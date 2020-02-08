package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const maxUploadSize = 1 << 20 // 1 MB

func upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(maxUploadSize)
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

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
