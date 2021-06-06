package main

import (
	"log"
	"net/http"
)

const authToken = "X-Auth-Token"

func requireAuthorizationMiddleware(al AuthLayerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(authToken)

			if !al.IsValidToken(Token(token)) {
				log.Printf("error while looking up session token: %s\n", token)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
