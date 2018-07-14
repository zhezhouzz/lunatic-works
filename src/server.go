package main

import (
    //"fmt"
    "log"
    "net/http"
//	"html/template"
)

func main() {
    err := http.ListenAndServe(":12754", http.FileServer(http.Dir("/Users/zhouzhe/go/src/github.com/zhezhouzz/lunatic-works/src/web")))
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
