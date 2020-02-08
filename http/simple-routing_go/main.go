package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
}

func aboutus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "We are friendly folk!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/aboutus", aboutus)
	log.Fatal(http.ListenAndServe(":80", r))
}
