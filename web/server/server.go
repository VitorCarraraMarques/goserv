package server

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/websocket"

	"github.com/VitorCarraraMarques/goserv/pkg/subway"
	"github.com/VitorCarraraMarques/goserv/web/server/auth"
	"github.com/VitorCarraraMarques/goserv/web/server/data"
)

const ADDR = "127.0.0.1"
const PORT = "8000"

type SubData struct {
	Nodes    []subway.Node
	Lanes    []subway.Lane
	LanesMap map[int]subway.Lane
	Paths    [][]subway.Node
	Add      func(int, int) int
}

type ContextData struct {
	Location      string
	ProjectData   []data.Project
	DraggableData []data.Draggable
	Subway        SubData
}

func Serve() {
	auth.LoadEnvVars()

	//Templates
	http.HandleFunc("/drag", draghandler)
	http.HandleFunc("/projects", projectshandler)
	http.HandleFunc("/message", messagehandler)
	http.HandleFunc(fmt.Sprintf("/subway/admin/%s", os.Getenv("SUBWAY_ADMIN_PATH")), subwayadminhandler)
	http.HandleFunc("/subway/admin", subwayloginhandler)
	http.HandleFunc("/subway/add", subwayaddhandler)
	http.HandleFunc("/subway/nodes", subwaylistnodeshandler)
	http.HandleFunc("/subway/path", subwaypathhandler)
	http.HandleFunc("/subway", subwayhandler)
	http.HandleFunc("/", basichandler)

	//WebSockets
	http.Handle("/echo", websocket.Handler(EchoServer))

	//Static
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/client/template/assets/"))))
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("web/client/script/"))))
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("web/client/src/"))))
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("web/client/style/"))))

	//Initial Log
	url := fmt.Sprintf("%v:%v", ADDR, PORT)
	fmt.Printf("[INFO] Listening on %v...\n", url)

	//Serve
	log.Fatal("[CRITICAL] Server failed :\n", http.ListenAndServe(url, nil))
}

func basichandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	base := template.Must(template.ParseFiles(
		"web/client/template/base.html",
		"web/client/template/header.html",
		"web/client/template/nav.html",
		"web/client/template/default.html",
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
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	projects := template.Must(template.ParseFiles(
		"web/client/template/base.html",
		"web/client/template/header.html",
		"web/client/template/nav.html",
		"web/client/template/projects.html",
	))
	path := r.URL.Path
	context := ContextData{}
	context.Location = path
	context.ProjectData = append(context.ProjectData, data.MyProjects...)
	err := projects.Execute(w, context)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func messagehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(
		"web/client/template/message.html",
	))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func draghandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(
		"web/client/template/drag.html",
	))
	context := ContextData{}
	context.DraggableData = append(context.DraggableData, data.MyDraggable...)
	err := tmpl.Execute(w, context)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func subwayhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(
		"web/client/template/subway/main.html",
	))

	lanes := subway.ListAllLanes()
	sd := SubData{
		Lanes: lanes,
	}
	ctx := ContextData{
		Subway: sd,
	}

	err := tmpl.Execute(w, ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func subwaylistnodeshandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(
		"web/client/template/subway/stationsdropdown.html",
	))

	params := r.URL.Query()
	l, err := strconv.Atoi(params.Get("SelectedLane"))
	if err != nil {
		log.Fatalf("[ERROR]Failed reading lane id: %s", err.Error())
	}
	nodes := subway.FindNodesByLane(l)
	s := SubData{
		Nodes: nodes,
	}
	ctx := ContextData{
		Subway: s,
	}

	err = tmpl.Execute(w, ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func subwaypathhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	err := r.ParseForm()
	if err != nil {
		log.Fatal("[ERROR] Failed at Parsing Form")
	}

	src, err := strconv.Atoi(r.Form.Get("StationFrom"))
	if err != nil {
		log.Fatalf("[ERROR]Invalid Source: %s", err.Error())
	}
	dest, err := strconv.Atoi(r.Form.Get("StationTo"))
	if err != nil {
		log.Fatalf("[ERROR]Invalid Destination: %s", err.Error())
	}

	lanes := subway.MapAllLanes()
	paths := subway.FindPaths(src, dest)

    f := [][]subway.Node{paths[0]}
	ctx := SubData{
		LanesMap: lanes,
		Paths:    f,
		Add:      func(a int, b int) int { return a + b },
	}
	tmpl := template.Must(template.ParseFiles(
		"web/client/template/subway/path.html",
	))
	err = tmpl.Execute(w, ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func subwayaddhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)

	err := r.ParseForm()
	if err != nil {
		log.Fatal("[ERROR] Failed at Parsing Form")
	}
	name := r.PostForm.Get("NewNode")
	node := subway.Node{
		Name: name,
	}
	_, err = subway.InsertOneNode(node)
	if err != nil {
		log.Fatalf("[ERROR] Failed at saving node to DB: %s", err.Error())
	}

	tmpl := template.Must(template.ParseFiles(
		"web/client/template/subway/add.html",
	))

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func subwayadminhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(
		"web/client/template/subway/admin.html",
	))

	nodes := subway.ListAllNodes()
	err := tmpl.Execute(w, nodes)
	if err != nil {
		log.Fatalf("[ERROR]: %s", err.Error())
	}
}

func subwayloginhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles(
			"web/client/template/subway/login.html",
		))
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("[ERROR] Failed at reading Body")
		}

		usr_pwd := r.PostForm.Get("SubwayPassword")
		the_pwd := os.Getenv("SUBWAY_PASSWORD")

		if usr_pwd == the_pwd {
			tmpl := template.Must(template.ParseFiles(
				"web/client/template/subway/main.html",
			))
			admpath := os.Getenv("SUBWAY_ADMIN_PATH")
			resheader := w.Header()
			resheader.Set("HX-Redirect", fmt.Sprintf("/subway/admin/%s", admpath))
			err = tmpl.Execute(w, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			tmpl := template.Must(template.ParseFiles(
				"web/client/template/subway/forbidden.html",
			))
			err = tmpl.Execute(w, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
