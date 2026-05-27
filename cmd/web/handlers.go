package main 

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	"log"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Server","Go")


	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil{
		log.Print(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w,nil)

	if err != nil{
		log.Print(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request){
	id,err := 	strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}
	_,err = fmt.Fprintf(w,"Display a specific snippet with ID %d",id)
	if err != nil{
		fmt.Println("Write Error")
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	_,err := w.Write([]byte("Display a form for creating a new snippet..."))
	if err != nil{
		fmt.Println("Write Error")
	}
}

func snippetCreatePost(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusCreated)
	_,err := w.Write([]byte("Save a new snippet..."))
	if err != nil{
		fmt.Println("Write Error")
	}

}
