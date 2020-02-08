package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// The amount of memory to use before flushing to temporary files.
const maxMemory = 1 << 20 // 1 MB

// If you instead want to limit the size of requests handled, see
// https://godoc.org/net/http#MaxBytesReader

func upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		log.Print(err)
		return
	}

	data, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Print(err)
		return
	}
	defer data.Close()

	f, err := os.OpenFile("./tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()
	io.Copy(f, data)
}

func main() {
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
