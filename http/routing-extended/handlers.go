package main

import (
	"encoding/json"
	"net/http"
)

func handleSecretDataEndpoint() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("super secret stuff"))
	}
}

type SignInRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}
type SignInResponse struct {
	Token string `json:"token"`
}
func handleSignInEndpoint(auth AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		signInRequest := &SignInRequest{}

		if err := json.NewDecoder(r.Body).Decode(signInRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token := auth.SignIn(Login(signInRequest.Login), signInRequest.Password)
		if token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		response := SignInResponse{string(*token)}
		responseJSON, err := json.Marshal(&response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseJSON))
	}
}
