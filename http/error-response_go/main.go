package main

import (
	"fmt"
	"log"
	"net/http"
)

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, message)
}

func test400ClientError(w http.ResponseWriter, r *http.Request) {
	renderError(w, "test error", http.StatusBadRequest) // 400 is a generic "client error"
}

func test500ServerError(w http.ResponseWriter, r *http.Request) {
	renderError(w, "test server error", http.StatusInternalServerError) // 500 is a generic "server error"
}

func main() {
	http.HandleFunc("/client-error-path", test400ClientError)
	http.HandleFunc("/server-error-path", test500ServerError)
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
