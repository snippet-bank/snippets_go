package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, name string, value string) {
	expire := time.Now().AddDate(0, 0, 30) // 30 days
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func home(w http.ResponseWriter, r *http.Request) {
	setCookie(w, "visited", "1")
	fmt.Fprint(w, "i'm going to write a cookie...")
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
