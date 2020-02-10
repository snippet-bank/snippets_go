package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func notfound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sorry can't find that page!")
}

func main() {
	r := mux.NewRouter()
	var h http.Handler = http.HandlerFunc(notfound)
	r.NotFoundHandler = h

	log.Fatal(http.ListenAndServe(":80", r))
}
