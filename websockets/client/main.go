package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func connectToServer() *websocket.Conn {
	log.Println("connecting")
	connection, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/socket", nil)
	if err != nil {
		log.Fatalf("dialing error: %s\n", err)
	}
	return connection
}

func handleReceivingMessages(serverConnection *websocket.Conn) {
	for {
		_, message, err := serverConnection.ReadMessage()
		if err != nil {
			log.Println("reading error:", err)
			return
		}
		log.Printf("received message: %s", message)
	}
}

func handleInterrupt(serverConnection *websocket.Conn, serverClosedConnectionChannel chan struct{}) {
	log.Println("interrupted")

	// Cleanly close the serverConnection by sending a close message and then
	// waiting (with timeout) for the server to close the serverConnection.
	err := serverConnection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Printf("error while sending close message: %s\n", err)
		return
	}

	select {
	case <-serverClosedConnectionChannel:
	case <-time.After(time.Second):
	}
	return
}

func sendTimeToServer(serverConnection *websocket.Conn, time time.Time) error {
	err := serverConnection.WriteMessage(websocket.TextMessage, []byte(time.String()))
	if err != nil {
		log.Printf("sending message error: %s\n", err)
		return err
	}
	return nil
}

func main() {
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)

	serverConnection := connectToServer()
	defer serverConnection.Close()

	serverClosedConnectionChannel := make(chan struct{})
	defer close(serverClosedConnectionChannel)
	go handleReceivingMessages(serverConnection)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-serverClosedConnectionChannel:
			log.Println("server closed connection")
			return
		case time := <-ticker.C:
			log.Println("sending a message with current time")
			err := sendTimeToServer(serverConnection, time)
			if err != nil {
				return
			}
		case <-interruptChannel:
			log.Println("client interrupted")
			handleInterrupt(serverConnection, serverClosedConnectionChannel)
			return
		}
	}
}
