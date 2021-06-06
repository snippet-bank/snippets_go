package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handleSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrader error: %s\n", err)
		return
	}
	defer c.Close()

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("reading error: %s\n", err)
			return
		}

		log.Printf("received msg: %s\n", message)

		if messageType != websocket.TextMessage {
			log.Println("received msg with binary type")
			return
		}

		err = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Hi, you sent me message: %s", message)))

		if err != nil {
			log.Printf("writing error: %s\n", err)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/socket", handleSocket)

	log.Println("listening")
	log.Fatal(http.ListenAndServe(":3000", r))
}
