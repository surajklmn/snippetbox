package main 

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"

	"github.com/surajklmn/snippetbox/internal/models"
)
// home is now a methond against *apllication
func (app *application)home(w http.ResponseWriter, r *http.Request){
	// we added a middleware to write custom header so no need to do this
	//w.Header().Add("Server","Go")
	snippets, err := app.snippets.Latest()
	if err!= nil{
		app.serverError(w,r,err)
		return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	// New render helper
	app.render(w,r,http.StatusOK, "home.tmpl.html",data)

	// for _,snippet := range snippets{
	// 	fmt.Fprintf(w,"%v\n",snippet)
	// }


	// Initialize a slice containing the path to the two files.
	// Our base template must be the first file in the slice
// 	files := []string{
// 		"./ui/html/base.tmpl.html",
// 		"./ui/html/pages/home.tmpl.html",
// 		"./ui/html/partials/nav.tmpl.html",
// 	}
// // Parse the template files
// 	ts, err := template.ParseFiles(files...)
// 	if err != nil{
// 		app.serverError(w,r,err)
// 		return
// 	}
// 	data := templateData {
// 	Snippets : snippets,
// 	}
// 	//Execute template
// 	//Any data that you pass as the final parameter to ts.ExecuteTemplate() is represented
// // within your HTML templates by the . character (referred to as dot).
// 	err = ts.ExecuteTemplate(w,"base",data)
//
// 	if err != nil{
// 		app.serverError(w,r,err)
// 	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request){
	id,err := 	strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	snippet,err := app.snippets.Get(id)
	if err != nil{
		if errors.Is(err,models.ErrNoRecord){
			http.NotFound(w,r)
		}else{
			app.serverError(w,r,err)
		}
		return
	}
	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// 	"./ui/html/pages/view.tmpl.html",
	// }
	//
	// ts,err := template.ParseFiles(files...)
	// if err != nil{
	// 	app.serverError(w,r,err)
	// 	return
	// }
	//
	// data := templateData{
	// 	Snippet: snippet,
	// }
	//
	// err = ts.ExecuteTemplate(w,"base",data)
	// if err != nil{
	// 	app.serverError(w,r,err)
	// }

	//User the new render helper
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w,r,http.StatusOK,"view.tmpl.html",data)
}

func (app *application)snippetCreate(w http.ResponseWriter, r *http.Request){
	_,err := w.Write([]byte("Display a form for creating a new snippet..."))
	if err != nil{
		fmt.Println("Write Error")
	}
}

func (app *application)snippetCreatePost(w http.ResponseWriter,r *http.Request){
	title := "O snail"
	content := "O snail Climb Moun Fuji,jkdfjkajdkj"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil{
		app.serverError(w,r,err)
		return
	}

	http.Redirect(w,r,fmt.Sprintf("/snippet/view/%d",id),http.StatusSeeOther)


}
