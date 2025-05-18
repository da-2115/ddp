package components

// This is the file for the events table on the 'View-Scores' page uses htmx along with the table.html
// Each handler is one layer deeper than the last on the scores.html page

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/da-2115/ddp/web/auth"
	"github.com/da-2115/ddp/web/data"
	"github.com/da-2115/ddp/web/util"
)

var tablesTmpl *template.Template

func init() {
	tablesTmpl = util.Unwrap(template.ParseFiles("views/tables.html"))
}

func ScoresHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	c, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	id := auth.SessionMap[c.Value]

	page := r.URL.Query().Get("page")
	qType := r.URL.Query().Get("qType")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	switch qType {
	case "event":
		eventsListHandler(w, q, pageNum, id.ArcheryAustraliaId)
	case "round":
		roundListHandler(w, r, q, pageNum, id.ArcheryAustraliaId)
	case "range":
		rangeListHandler(w, r, q, pageNum, id.ArcheryAustraliaId)
	case "end":
		endListHandler(w, r, q, pageNum, id.ArcheryAustraliaId)
	case "score":
		scoreListHandler(w, r, q, pageNum, id.ArcheryAustraliaId)
	default:
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func eventsListHandler(w http.ResponseWriter, q *data.Queries, pageNum int, id string) {
	events, err := q.GetEventsByID(context.Background(), data.GetEventsByIDParams{
		Archeryaustraliaid: id,
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = tablesTmpl.ExecuteTemplate(w, "event-table", events)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func roundListHandler(w http.ResponseWriter, r *http.Request, q *data.Queries, pageNum int, id string) {
	eventID := r.URL.Query().Get("event_id")
	eventIDNum, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rounds, err := q.GetRoundsByID(context.Background(), data.GetRoundsByIDParams{
		Archeryaustraliaid: id,
		Eventid:            int32(eventIDNum),
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	type roundWrapper struct {
		Rows []data.GetRoundsByIDRow
		Page int
	}

	err = tablesTmpl.ExecuteTemplate(w, "round-table", roundWrapper{
		Rows: rounds,
		Page: pageNum,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func rangeListHandler(w http.ResponseWriter, r *http.Request, q *data.Queries, pageNum int, id string) {
	eventID := r.URL.Query().Get("event_id")
	eventIDNum, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	roundID := r.URL.Query().Get("round_id")
	roundIDNum, err := strconv.Atoi(roundID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ranges, err := q.GetRangesByID(context.Background(), data.GetRangesByIDParams{
		Archeryaustraliaid: id,
		Eventid:            int32(eventIDNum),
		Roundid:            int32(roundIDNum),
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	type rangeWrapper struct {
		Rows    []data.GetRangesByIDRow
		Page    int
		Eventid int
	}

	err = tablesTmpl.ExecuteTemplate(w, "range-table", rangeWrapper{
		Rows:    ranges,
		Page:    pageNum,
		Eventid: eventIDNum,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func endListHandler(w http.ResponseWriter, r *http.Request, q *data.Queries, pageNum int, id string) {
	eventID := r.URL.Query().Get("event_id")
	eventIDNum, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	roundID := r.URL.Query().Get("round_id")
	roundIDNum, err := strconv.Atoi(roundID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rangeID := r.URL.Query().Get("range_id")
	rangeIDNum, err := strconv.Atoi(rangeID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ends, err := q.GetEndsByID(context.Background(), data.GetEndsByIDParams{
		Archeryaustraliaid: id,
		Eventid:            int32(eventIDNum),
		Roundid:            int32(roundIDNum),
		Rangeid:            int32(rangeIDNum),
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	type endWrapper struct {
		Rows    []data.GetEndsByIDRow
		Page    int
		Eventid int
		Roundid int
	}

	err = tablesTmpl.ExecuteTemplate(w, "end-table", endWrapper{
		Rows:    ends,
		Page:    pageNum,
		Eventid: eventIDNum,
		Roundid: roundIDNum,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func scoreListHandler(w http.ResponseWriter, r *http.Request, q *data.Queries, pageNum int, id string) {
	eventID := r.URL.Query().Get("event_id")
	eventIDNum, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	roundID := r.URL.Query().Get("round_id")
	roundIDNum, err := strconv.Atoi(roundID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rangeID := r.URL.Query().Get("range_id")
	rangeIDNum, err := strconv.Atoi(rangeID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	endID := r.URL.Query().Get("end_id")
	endIDNum, err := strconv.Atoi(endID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	scores, err := q.GetScoresByID(context.Background(), data.GetScoresByIDParams{
		Archeryaustraliaid: id,
		Eventid:            int32(eventIDNum),
		Roundid:            int32(roundIDNum),
		Rangeid:            int32(rangeIDNum),
		Endid:              int32(endIDNum),
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	type scoreWrapper struct {
		Rows    []data.GetScoresByIDRow
		Page    int
		Eventid int
		Roundid int
		Rangeid int
	}

	err = tablesTmpl.ExecuteTemplate(w, "score-table", scoreWrapper{
		Rows:    scores,
		Page:    pageNum,
		Eventid: eventIDNum,
		Roundid: roundIDNum,
		Rangeid: rangeIDNum,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
