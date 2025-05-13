package components

import (
	"html/template"
	"net/http"

	"github.com/da-2115/ddp/web/util"
)

var navTmpl *template.Template

func init() {
	navTmpl = util.Unwrap(template.ParseFiles("views/nav.html"))
}

func NavHandler(w http.ResponseWriter, r *http.Request) {
	type navData struct{}
	navTmpl.ExecuteTemplate(w, "nav", navData{})
}
