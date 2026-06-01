package main 
import (
	"log"
	"net/http"
	"flag"

)

func main(){
	mux := http.NewServeMux()

	addr := flag.String("addr",":4000","HTTP network address")
	flag.Parse()


	//Create a file server which servers files out of the ".ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//Use the mux.Handle fucntion to register the fileServer as the handler for all URL path that start with "/static/"
	mux.Handle("GET /static/",http.StripPrefix("/static",fileServer))


	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}",snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create/",snippetCreatePost)

	
	log.Printf("starting server  on %s",*addr)
	err := http.ListenAndServe(*addr,mux)
	log.Fatal(err)


}
