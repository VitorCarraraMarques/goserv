package server

import (
	"html/template"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/template/assets/"))))
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	base := template.Must(template.ParseFiles(
		"web/template/base.html",
		"web/template/header.html",
		"web/template/nav.html",
	))
    path := r.URL.Path
    context := make(map[string]string)
    context["Location"] = path
	base.Execute(w, context)
}
