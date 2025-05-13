package auth

import (
	"net/http"
	"time"
)

type Session struct {
	ArcheryAustraliaId string
	Admin              bool
	Expires            time.Time
}

// this is where we store the cookies on the server
var SessionMap map[string]Session

func init() {
	SessionMap = make(map[string]Session)
}

// checks if the cookie exists on the server based on the `Expires` time.Time value
func cookieIsValid(c *http.Cookie) bool {
	if s, exists := SessionMap[c.Value]; exists {
		if time.Now().Before(s.Expires) {
			return true
		} else {
			delete(SessionMap, c.Value)
			return false
		}
	}
	return false
}

// auth middleware that checks to see if the user has a session valid cookie, kicks them to the login screen if they don't
func AuthMiddleware(next http.Handler) http.Handler {
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

// Checks if requester is an admin (club recorder), expects AuthMiddleware to have already validated
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if s, exists := SessionMap[c.Value]; exists {
			if !s.Admin {
				http.Redirect(w, r, "/"+r.URL.Path[1:], http.StatusSeeOther)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// this is just a test handler for the button on the top right of the screen
func AuthTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// as this is after ran after AuthMiddleware we know that the user is logged in
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	if s, exists := SessionMap[cookie.Value]; exists {
		if s.Admin {
			w.Write([]byte("You are authenticated\nYou are a club recorder"))
			return
		}
	}
	w.Write([]byte("You are authenticated"))
}
