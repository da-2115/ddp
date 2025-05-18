package components

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/da-2115/ddp/web/data"
	"github.com/da-2115/ddp/web/util"
)

var listTmpl *template.Template

func init() {
	listTmpl = util.Unwrap(template.ParseFiles("views/list.html"))
}

func ViewAllEventsHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var page int
	var err error
	if page, err = strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	type wrapper struct {
		Rows []data.Event
		Page int
	}

	events, err := q.GetAllEvents(context.Background(), data.GetAllEventsParams{
		Limit:  10,
		Offset: (int32(page) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = listTmpl.ExecuteTemplate(w, "events", wrapper{Page: page, Rows: events})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func ViewAllRoundsHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var page int
	var err error
	if page, err = strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var eventID int
	if eventID, err = strconv.Atoi(r.URL.Query().Get("event_id")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	type wrapper struct {
		Rows []data.Round
		Page int
	}

	rounds, err := q.GetRoundByEvent(context.Background(), data.GetRoundByEventParams{
		Eventid: int32(eventID),
		Limit:   10,
		Offset:  (int32(page) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = listTmpl.ExecuteTemplate(w, "rounds", wrapper{Page: page, Rows: rounds})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func ViewAllRangesHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var page int
	var err error
	if page, err = strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var roundID int
	if roundID, err = strconv.Atoi(r.URL.Query().Get("round_id")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	type wrapper struct {
		Rows    []data.Range
		Page    int
	}

	ranges, err := q.GetRangeByRound(context.Background(), data.GetRangeByRoundParams{
		Roundid: int32(roundID),
		Limit:   10,
		Offset:  (int32(page) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = listTmpl.ExecuteTemplate(w, "ranges", wrapper{Page: page, Rows: ranges})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func ViewAllEndsHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var page int
	var err error
	if page, err = strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var rangeID int
	if rangeID, err = strconv.Atoi(r.URL.Query().Get("range_id")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	type wrapper struct {
		Rows    []data.End
		Page    int
	}

	ends, err := q.GetEndByRound(context.Background(), data.GetEndByRoundParams{
		Rangeid: int32(rangeID),
		Limit:   10,
		Offset:  (int32(page) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = listTmpl.ExecuteTemplate(w, "ends", wrapper{Page: page, Rows: ends})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func ViewAllScoresHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var page int
	var err error
	if page, err = strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var endID int
	if endID, err = strconv.Atoi(r.URL.Query().Get("end_id")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	type wrapper struct {
		Rows  []data.Score
		Page  int
	}

	scores, err := q.GetScoreByRound(context.Background(), int32(endID))

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = listTmpl.ExecuteTemplate(w, "scores", wrapper{Page: page, Rows: scores})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
