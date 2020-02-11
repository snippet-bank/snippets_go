package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var delaySecs int

func init() {
	delaySecs, err := strconv.Atoi(os.Getenv("DELAY_SECS"))
	if err != nil || delaySecs == 0 {
		delaySecs = 10
	}
	log.Printf("Delay set to %d secs", delaySecs)
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/delayme", slowHandler)

	log.Println("Running slow server...")
	log.Fatal(http.ListenAndServe(":3000", http.DefaultServeMux))
}
