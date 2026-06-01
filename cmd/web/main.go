package main 
import (
	"log/slog"
	"net/http"
	"flag"
	"os"

)

type application struct{//application struct to hold system wide dependencies
	logger *slog.Logger
}

func main(){
	mux := http.NewServeMux()

	addr := flag.String("addr",":4000","HTTP network address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{
		AddSource :true,
	}))

	app := &application{
		logger: logger,
	}

	//Create a file server which servers files out of the ".ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//Use the mux.Handle fucntion to register the fileServer as the handler for all URL path that start with "/static/"
	mux.Handle("GET /static/",http.StripPrefix("/static",fileServer))


	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}",app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create/",app.snippetCreatePost)

	logger.Info("starting server",slog.String("addr", *addr))
	
	err := http.ListenAndServe(*addr,mux)

	logger.Error(err.Error())
	os.Exit(1)

}
