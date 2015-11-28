package main

import (
	"encoding/json"
	"fmt"
	"github.com/gpeden/hello-go/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

type (
	Quote struct {
		Quote string `json:"quote"`
		Id    string `json:"id"`
	}
)

func main() {
	router := httprouter.New()
	router.GET("/", HelloHandler)
	router.GET("/quote", QuoteHandler)
	fmt.Println("listening...")

	http.Handle("/", router)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func HelloHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(res, "go: HelloHandler: hello, world!!!")
}

func QuoteHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// Stub an example quote
	q := Quote{
		Quote: "hey this is a private residence man!",
		Id:    p.ByName("id"),
	}

	qj, _ := json.Marshal(q)

	// Write content-type, statuscode, payload
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	fmt.Fprintf(res, "%s", qj)
}
