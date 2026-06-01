package main 
import (
	"log/slog"
	"net/http"
	"flag"
	"os"

)

func main(){
	mux := http.NewServeMux()

	addr := flag.String("addr",":4000","HTTP network address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{
		AddSource :true,
	}))

	//Create a file server which servers files out of the ".ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//Use the mux.Handle fucntion to register the fileServer as the handler for all URL path that start with "/static/"
	mux.Handle("GET /static/",http.StripPrefix("/static",fileServer))


	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}",snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create/",snippetCreatePost)

	logger.Info("starting server",slog.String("addr", *addr))
	
	err := http.ListenAndServe(*addr,mux)

	logger.Error(err.Error())
	os.Exit(1)

}
