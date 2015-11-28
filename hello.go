package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", hello)
    router.HandleFunc("/quote", quote)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "go: hello, world!!!")
}

func quote(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hey this is a private residence man!")
}
