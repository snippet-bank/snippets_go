package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const staticDir = "/static/"

func sampleAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"test\": true}")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/sample", sampleAPIHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":80", r))
}
