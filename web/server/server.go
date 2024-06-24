package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func Serve() {
    http.HandleFunc("/", home)
    http.HandleFunc("/greet", greet)
    http.HandleFunc("/curse", curse)
	http.ListenAndServe("localhost:8000", nil)
}


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("web/server/server.go::home()")
	tmpl, err := template.ParseFiles("web/template/base.tmpl", "web/template/home.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("web/server/server.go::greet()")
	tmpl, err := template.ParseFiles("web/template/base.tmpl", "web/template/greet.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}


func curse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("web/server/server.go::curse()")
	tmpl, err := template.ParseFiles("web/template/base.tmpl", "web/template/curse.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}
