package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
