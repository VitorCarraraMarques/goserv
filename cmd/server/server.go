package main

import (
    "fmt"
    "github.com/VitorCarraraMarques/goserv/web/server"
)

func main(){
    fmt.Println("THIS IS cmd/server/main")
    server.Serve()
}

