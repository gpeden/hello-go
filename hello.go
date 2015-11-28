package main

import (
	"encoding/json"
	"fmt"
	"github.com/gpeden/hello-go/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"net/http"
	"os"
  "strconv"
)


var quotes = []Quote{
    {Quote: "Just one thing, Dude. Do ya have to use so many cuss words?"},
    {Quote: "Hello. Mein dizbatcher says zere iss somezing wrong mit deine kable."},
    {Quote: "Vee vant zat money, Lebowski."},
    {Quote: "Uh, yeah. Probably a vagrant, slept in the car. Or perhaps just used it as a toilet and moved on."},
    {Quote: "Your name is Lebowski, Lebowski. Your wife is Bunny."},
    {Quote: "You are joking. But perhaps you are right."},
    {Quote: "I'm staying. I'm finishing my coffee. Enjoying my coffee."},
    {Quote: "Tomorrow vee come back and cut off your Johnson."},
    {Quote: "Has the whole world gone CRAZY?! Am I the only one who gives a shit about the rules?! MARK IT ZERO!"},

}

func main() {


	router := httprouter.New()
	router.GET("/", HelloHandler)
	router.GET("/quote/:id", QuoteHandler)
  router.GET("/quotes", QuotesHandler)
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
	// This id business seems a little hacky
  id, _ := strconv.Atoi(p.ByName("id"))
	q := quotes[id - 1]

  // Set the ID in the return object
  q.Id = p.ByName("id")

	qj, _ := json.Marshal(q)

	// Write content-type, statuscode, payload
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	fmt.Fprintf(res, "%s", qj)
}

func QuotesHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	qj, _ := json.Marshal(quotes)

	// Write content-type, statuscode, payload
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	fmt.Fprintf(res, "%s", qj)
}
