package components

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/da-2115/ddp/web/data"
	"github.com/da-2115/ddp/web/util"
)

var stageTmpl *template.Template

func init() {
	stageTmpl = util.Unwrap(template.ParseFiles("views/stage-list.html"))
}

func StageEndListHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	ends, err := q.GetStagedEnds(context.Background())
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	err = stageTmpl.ExecuteTemplate(w, "stage", ends)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StageScoreHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var err error
	var endID int
	if endID, err = strconv.Atoi(r.URL.Query().Get("end_id")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	scores, err := q.GetScoreByEnd(context.Background(), int32(endID))
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	type wrapper struct {
		Rows  []data.Score
		Endid int
	}

	err = stageTmpl.ExecuteTemplate(w, "score", wrapper{Endid: endID, Rows: scores})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func StageEndHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var err error
	var endID int
	if endID, err = strconv.Atoi(r.FormValue("end_id")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = q.StageEnd(context.Background(), int32(endID))
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	StageEndListHandler(w,r,q)
}

func DeleteEndHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var endID int
	var err error
	if endID, err = strconv.Atoi(r.FormValue("end_id")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = q.DeleteEnd(context.Background(), int32(endID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	StageEndListHandler(w,r,q)
}
