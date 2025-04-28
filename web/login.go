package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"net/http"
	"time"

	"github.com/da-2115/ddp/web/data"
	"golang.org/x/crypto/bcrypt"
)

func loginHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")
	next := r.FormValue("next")

	m, err := q.GetMemberByID(context.Background(), user)
	if err != nil {
		slog.Info("Login Request Invalid User", "user", user)
		http.Redirect(w, r, "/login.html?next=" + next, http.StatusSeeOther)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.Passwordhash), []byte(pass)); err == nil {

		b := make([]byte, 32)
		_, err := rand.Read(b)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		cookieVal := base64.URLEncoding.EncodeToString(b)
		sessionMap[cookieVal] = session{
			ArcheryAustraliaId: user,
			Expires:            time.Now().Add(24 * time.Hour),
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    cookieVal,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   false, // https
			SameSite: http.SameSiteLaxMode,
		})

		slog.Info("Login success", "user", user, "next", next)
		http.Redirect(w, r, "/" + next, http.StatusSeeOther)
		return
	} else {
		slog.Info("Login wrong password", "user", user)
		http.Redirect(w, r, "/login.html?next=" + next, http.StatusSeeOther)
		return
	}
}
