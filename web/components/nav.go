package components

import (
	"html/template"
	"net/http"

	"github.com/da-2115/ddp/web/auth"
	"github.com/da-2115/ddp/web/util"
)

var navTmpl *template.Template

func init() {
	navTmpl = util.Unwrap(template.ParseFiles("views/nav.html"))
}

func NavHandler(w http.ResponseWriter, r *http.Request) {
	var admin bool

	cookie, err := r.Cookie("session_id")
	if err != nil && err != http.ErrNoCookie {
		http.Error(w, "", http.StatusBadRequest)
		return
	} else {
		if cookie != nil {
			s, exists := auth.SessionMap[cookie.Value]
			if exists {
				admin = s.Admin
			}
		}
	}

	navTmpl.ExecuteTemplate(w, "nav", admin)
}
