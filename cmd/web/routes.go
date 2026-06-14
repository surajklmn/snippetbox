package main 

import "net/http"

func (app *application) routes() http.Handler{
	mux := http.NewServeMux()

	//Create a file server which servers files out of the ".ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))

	//Use the mux.Handle fucntion to register the fileServer as the handler for all URL path that start with "/static/"
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))

	mux.HandleFunc("GET /{$}",app.home)
	mux.HandleFunc("GET /snippet/view/{id}",app.snippetView)
	mux.HandleFunc("GET /snippet/create",app.snippetCreate)
	mux.HandleFunc("POST /snippet/create",app.snippetCreatePost)

	return app.logRequest(commonHeaders(mux))
}
