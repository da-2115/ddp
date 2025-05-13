package main

import (
	"net/http"
	"time"
)

type session struct {
	ArcheryAustraliaId string
	Expires            time.Time
}

// this is where we store the cookies on the server
var sessionMap map[string]session

func init() {
	sessionMap = make(map[string]session)
}

// checks if the cookie exists on the server based on the `Expires` time.Time value
func cookieIsValid(c *http.Cookie) bool {
	if s, exists := sessionMap[c.Value]; exists {
		if time.Now().Before(s.Expires) {
			return true
		} else {
			delete(sessionMap, c.Value)
			return false
		}
	}
	return false
}

// auth middleware that checks to see if the user has a session valid cookie, kicks them to the login screen if they don't
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/login.html?next="+r.URL.Path[1:], http.StatusSeeOther)
			return
		}

		if !cookieIsValid(cookie) {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_id",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),
				MaxAge:  -1,
			})
			// ?next is used to redirect back after logging in
			http.Redirect(w, r, "/login.html?next="+r.URL.Path[1:], http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// this is just a test handler for the button on the top right of the screen
func authTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("You are authenticated"))
}
