package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Serve() {
	fmt.Println("[[[ Listening on port 8000 ]]]")
	http.HandleFunc("/", basichandler)
    http.HandleFunc("/projects", projectshandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/template/assets/"))))
	http.ListenAndServe("localhost:8000", nil)
}

type ContextData struct {
	Location    string
	ProjectData []Project
}

func basichandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n  >>> %s %s \n", r.Method, r.URL.Path)
	base := template.Must(template.ParseFiles(
		"web/template/base.html",
		"web/template/header.html",
		"web/template/nav.html",
	))
	path := r.URL.Path
	context := ContextData{}
	context.Location = path
	err := base.Execute(w, context)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func projectshandler(w http.ResponseWriter, r *http.Request) {
	base := template.Must(template.ParseFiles(
		"web/template/base.html",
		"web/template/header.html",
		"web/template/nav.html",
	))
	path := r.URL.Path
	context := ContextData{}
	context.Location = path
    context.ProjectData = append(context.ProjectData, MyProjects...)
    err := base.Execute(w, context)
    if err != nil {
		log.Fatal(err.Error())
    }
}
