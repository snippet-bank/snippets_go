package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi")
}

func currentTime(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

// Logger is a middleware for request logging
type Logger struct {
	handler http.Handler
}

// ServeHTTP method on the Logger struct logs the request details as a side effect.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handler http.Handler) *Logger {
	// wrap the real handler
	return &Logger{handler}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/time", currentTime)
	wrappedMux := NewLogger(mux)

	log.Printf("server is listening")
	log.Fatal(http.ListenAndServe(":80", wrappedMux))
}
