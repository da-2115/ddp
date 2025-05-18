package components

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/da-2115/ddp/web/auth"
	"github.com/da-2115/ddp/web/data"
	"github.com/da-2115/ddp/web/util"
)

var submitTmpl *template.Template

func init() {
	submitTmpl = util.Unwrap(template.ParseFiles("views/events-list.html"))
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {

}

func SubmitEventsHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {

	page := r.URL.Query().Get("page")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	events, err := q.GetEvents(context.Background(), data.GetEventsParams{
		Limit:  10,
		Offset: (int32(pageNum) - 1) * 10,
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = submitTmpl.ExecuteTemplate(w, "event-list", events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SubmitRoundsHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	page := r.URL.Query().Get("page")
	cookie, err := r.Cookie("session_id")

	s, exists := auth.SessionMap[cookie.Value]
	if !exists {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}


	eventID := r.URL.Query().Get("event_id")
	eventIDNum, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rounds, err := q.GetRounds(context.Background(), data.GetRoundsParams{
		Eventid: int32(eventIDNum),
		Archeryaustraliaid: s.ArcheryAustraliaId,
		Limit:  10,
		Offset: (int32(pageNum) - 1) * 10,
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	type RoundsWrapper struct {
		Rows []data.GetRoundsRow
		Page int
	}

	err = submitTmpl.ExecuteTemplate(w, "round-list", RoundsWrapper{Rows: rounds, Page: pageNum})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func SubmitRangesHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	page := r.URL.Query().Get("page")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	roundID := r.URL.Query().Get("round_id")
	roundIDNum, err := strconv.Atoi(roundID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ranges, err := q.GetRangesByRound(context.Background(), data.GetRangesByRoundParams{
		Roundid: int32(roundIDNum),
		Limit: 10,
		Offset: (int32(pageNum) - 1) * 10,
	})

	type RangesWrapper struct {
		Rows []data.Range
		Eventid string
		Page int
	}

	eventID := r.URL.Query().Get("event_id")

	obj := RangesWrapper{Rows: ranges, Page: pageNum, Eventid: eventID}

	err = submitTmpl.ExecuteTemplate(w, "range-list", obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
