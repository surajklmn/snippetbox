package main 
import (
	"database/sql"
	"log/slog"
	"net/http"
	"flag"
	"os"
	"github.com/surajklmn/snippetbox/internal/models"
	_"github.com/go-sql-driver/mysql"

)

type application struct{//application struct to hold system wide dependencies
	logger *slog.Logger
	snippets *models.SnippetModel
}

func main(){

	addr := flag.String("addr",":4000","HTTP network address")
	dsn := flag.String("dsn","web:pass@/snippetbox?parseTime=true","MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{
		AddSource :true,
	}))
	db, err := openDB(*dsn)
	if err != nil{
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB:db},
	}

	logger.Info("starting server",slog.String("addr", *addr))

	//Call the app.routes() method and pass that to ListenAndServe
	err = http.ListenAndServe(*addr,app.routes())

	logger.Error(err.Error())
	os.Exit(1)

}

func openDB(dsn string)(*sql.DB, error){
	db,err := sql.Open("mysql",dsn)
	if err!= nil{
		return nil, err
	}

	err = db.Ping()
	if err != nil{
		db.Close()
		return nil,err
	}
	return db, nil
}
