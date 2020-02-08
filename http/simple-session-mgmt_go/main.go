package main

import (
	"fmt"
	"log"
	"net/http"
)

func doSessionStuff(w http.ResponseWriter, r *http.Request) SessionID {
	// This is just a simple demo of the most-basic mechanisms of session management.
	// In a real app this session processing would be added to all handlers automatically
	// via a "middleware" (a function that looks at all requests and possibly modifies them).
	sid := getSessionIDFromCookie(r)
	if sid == "" {
		sid = sessionCreateAndSetCookie(w, r)
		log.Printf("No session ID found in cookie. Created session with id %s", sid)
	} else {
		log.Printf("Found existing session with id %s", sid)
	}
	return sid
}

func home(w http.ResponseWriter, r *http.Request) {
	doSessionStuff(w, r)
	fmt.Fprint(w, "welcome home!")
}

func someRoute(w http.ResponseWriter, r *http.Request) {
	sid := doSessionStuff(w, r)
	sessionWrite(sid, "test-key", "test-value")
	log.Printf("added session key/value pair: 'test-key=test-value'")

	fmt.Fprintf(w, "some route in the app...")
}

func anotherRoute(w http.ResponseWriter, r *http.Request) {
	sid := doSessionStuff(w, r)
	val, err := sessionRead(sid, "test-key")
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Printf("read value for key 'test-key' in session: %s", val)

	fmt.Fprint(w, "another route in the app...")
}

func logout(w http.ResponseWriter, r *http.Request) {
	sid := getSessionIDFromCookie(r)
	if sid != "" {
		sessionDestroyAndRemoveCookie(w, r, sid)
		log.Printf("nuked the session!")
	}
	fmt.Fprint(w, "Au revoir!")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/some-route", someRoute)
	http.HandleFunc("/another-route", anotherRoute)
	http.HandleFunc("/logout", logout)
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
