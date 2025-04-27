package main

import (
	"net/http"
	"time"
)

type session struct {
	ArcheryAustraliaId string
	Expires            time.Time
}

var sessionMap map[string]session

func init() {
	sessionMap = make(map[string]session)
}

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

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/login.html", http.StatusSeeOther)
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
			http.Redirect(w, r, "/login.html", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("You are authenticated"))
}
