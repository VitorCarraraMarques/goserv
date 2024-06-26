 package server 

 import ( 
     "html/template" 
     "log" 
     "net/http"
 )

type Project struct { 
     name string 
} 

const oneProject = Project{"Project"}

const MyProjects = []Project{ 
    Project(name="Carrara Software"),
    {name="LandingPageV1"}, 
    {name="EhPrimo"}
}
func Serve() {
	http.HandleFunc("/", home)
	http.HandleFunc("/projects", projects)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/template/assets/"))))
	http.ListenAndServe("localhost:8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"web/template/base.tmpl",
		"web/template/header.tmpl",
		"web/template/home.tmpl",
	)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func projects(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"web/template/base.html",
		"web/template/header.html",
		"web/template/projects.html",
	)
	if err != nil {
		log.Fatal(err)
	}
}





