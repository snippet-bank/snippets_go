package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	authLayer := NewAuthLayerImpl()
	err := Start(80, authLayer)
	if err != nil {
		log.Fatal(err)
	}
}

func Start(port int, al AuthLayerInterface) error {
	router := mux.NewRouter()
	attachHandlers(router, al)
	attachReqAuth(router, al)

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	log.Println("Server started...")
	log.Printf("Listening on %s\n", s.Addr)
	return s.ListenAndServe()
}

func attachHandlers(mux *mux.Router, al AuthLayerInterface) {
	mux.HandleFunc("/session", handleSignInEndpoint(al)).Methods(http.MethodPut)
}

func attachReqAuth(mux *mux.Router, al AuthLayerInterface) {
	muxAuth := mux.PathPrefix("").Subrouter()
	muxAuth.Use(requireAuthorizationMiddleware(al))

	muxAuth.HandleFunc("/secret", handleSecretDataEndpoint()).Methods(http.MethodGet)
}
