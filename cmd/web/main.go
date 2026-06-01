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

	addr := flag.String("addr",":4000","HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{
		AddSource :true,
	}))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server",slog.String("addr", *addr))

	//Call the app.routes() method and pass that to ListenAndServe
	err := http.ListenAndServe(*addr,app.routes())

	logger.Error(err.Error())
	os.Exit(1)

}
