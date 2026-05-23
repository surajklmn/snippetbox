package main 

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	_,err := w.Write([]byte("Hello from snippetbox"))

	if err != nil{
		log.Println("Write Error:",err)
	}
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", home)
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)

}

