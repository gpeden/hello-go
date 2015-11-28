package main

import (
	"fmt"
	"github.com/gpeden/hello-go/Godeps/_workspace/src/github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HelloHandler)
	router.HandleFunc("/quote", QuoteHandler)
	fmt.Println("listening...")

	http.Handle("/", router)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "go: HelloHandler: hello, world!!!")
}

func QuoteHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hey this is a private residence man!")
}
