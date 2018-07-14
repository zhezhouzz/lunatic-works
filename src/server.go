package main

import (
    "fmt"
    "log"
    "net/http"
	"html/template"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Inside HelloServer handler")
    fmt.Fprintf(w, "Welcome!\nThis is Lunatic Works!\n")
	t, err := template.ParseFiles("main.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServe(":12754", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
