package main 

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
)
// home is now a methond against *apllication
func (app *application)home(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Server","Go")

	// Initialize a slice containing the path to the two files.
	// Our base template must be the first file in the slice
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil{
		app.serverError(w,r,err)
		return
	}
	err = ts.ExecuteTemplate(w,"base",nil)

	if err != nil{
		app.serverError(w,r,err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request){
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

func (app *application)snippetCreate(w http.ResponseWriter, r *http.Request){
	_,err := w.Write([]byte("Display a form for creating a new snippet..."))
	if err != nil{
		fmt.Println("Write Error")
	}
}

func (app *application)snippetCreatePost(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusCreated)
	_,err := w.Write([]byte("Save a new snippet..."))
	if err != nil{
		fmt.Println("Write Error")
	}

}
