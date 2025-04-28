package main

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/da-2115/ddp/web/data"
)

var tmpl *template.Template

func init() {
	tmpl = unwrap(template.ParseFiles("./views/tables.html"))
}

func scoresHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	c, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	id := sessionMap[c.Value]

	page := r.URL.Query().Get("page")
	qType := r.URL.Query().Get("qType")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	switch qType {
	case "event":
		eventsListHandler(w, r, q, pageNum, id.ArcheryAustraliaId)
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
	}
}

func eventsListHandler(w http.ResponseWriter, r *http.Request, q *data.Queries, pageNum int, id string) {
	events, err := q.GetEvents(context.Background(), data.GetEventsParams{
		Archeryaustraliaid: id,
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
	})

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "event-table", events)
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

	rounds, err := q.GetRounds(context.Background(), data.GetRoundsParams{
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
		Rows []data.GetRoundsRow
		Page int
	}

	err = tmpl.ExecuteTemplate(w, "round-table", roundWrapper{
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

	ranges, err := q.GetRanges(context.Background(), data.GetRangesParams{
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
		Rows    []data.GetRangesRow
		Page    int
		Eventid int
	}

	err = tmpl.ExecuteTemplate(w, "range-table", rangeWrapper{
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
		http.Error(w, "", http.StatusBadGateway)
		return
	}

	rangeID := r.URL.Query().Get("range_id")
	rangeIDNum, err := strconv.Atoi(rangeID)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ends, err := q.GetEnds(context.Background(), data.GetEndsParams{
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
		Rows    []data.GetEndsRow
		Page    int
		Eventid int
		Roundid int
	}

	err = tmpl.ExecuteTemplate(w, "end-table", endWrapper{
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

	scores, err := q.GetScores(context.Background(), data.GetScoresParams{
		Archeryaustraliaid: id,
		Eventid:            int32(eventIDNum),
		Roundid:            int32(roundIDNum),
		Rangeid:            int32(rangeIDNum),
		Endid: 				int32(endIDNum),
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	type scoreWrapper struct {
		Rows    []data.GetScoresRow
		Page    int
		Eventid int
		Roundid int
		Rangeid int
	}

	err = tmpl.ExecuteTemplate(w, "score-table", scoreWrapper{
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
