package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const cookieName = "sid"
const cookieSecondsToExpiry = 3600

// Session is the server-side key-value "databag" for a given client session.
type Session map[string]string

// SessionID is the unique id of this session.
type SessionID string

var sessions = make(map[SessionID]Session) // sessions are stored in memory, in this simple implementation

func getSessionIDFromCookie(r *http.Request) SessionID {
	cookie, err := r.Cookie(cookieName)
	if err != nil || cookie.Value == "" {
		return ""
	}
	sid := SessionID(cookie.Value)
	return sid
}

func sessionCreateAndSetCookie(w http.ResponseWriter, r *http.Request) SessionID {
	var sid = SessionID(uuid.New().String())

	// populate session store
	sessions[sid] = Session{}

	// set session cookie on client
	cookie := http.Cookie{Name: cookieName, Value: string(sid), Path: "/", HttpOnly: true, MaxAge: cookieSecondsToExpiry}
	http.SetCookie(w, &cookie)

	return sid
}

func sessionRead(sid SessionID, key string) (string, error) {
	session, ok := sessions[sid]
	if !ok {
		return "", fmt.Errorf("session not found for sid %s", sid)
	}

	value, ok := session[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}

	return value, nil
}

func sessionWrite(sid SessionID, key string, value string) (Session, error) {
	session, ok := sessions[sid]
	if !ok {
		return nil, fmt.Errorf("session not found")
	}
	session[key] = value
	return session, nil
}

func sessionDestroyAndRemoveCookie(w http.ResponseWriter, r *http.Request, sid SessionID) {
	if _, ok := sessions[sid]; ok {
		delete(sessions, sid)
	}

	// remove session cookie on client
	expiredCookie := http.Cookie{Name: cookieName, Path: "/", HttpOnly: true, Expires: time.Now(), MaxAge: -1}
	http.SetCookie(w, &expiredCookie)
}
