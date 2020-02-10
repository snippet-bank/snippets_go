package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Here's where we set the response header.
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "<h1>HELLO</h1>")
	fmt.Fprint(w, "<p>foo</p>")
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
