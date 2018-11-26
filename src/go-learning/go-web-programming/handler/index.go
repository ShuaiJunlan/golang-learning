package handler

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := nil, nil
	templates.ExecuteTemplate(w, "layout", threads)

}
