package components

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/da-2115/ddp/web/auth"
	"github.com/da-2115/ddp/web/data"
	"github.com/da-2115/ddp/web/util"
)

var submitTmpl *template.Template
var scoresMap map[string]int32

func init() {
	submitTmpl = util.Unwrap(template.ParseFiles("views/submit/events-list.html", "views/submit/submit-form.html"))
	scoresMap = map[string]int32{
		"X":  10,
		"10": 10,
		"9":  9,
		"8":  8,
		"7":  7,
		"6":  6,
		"5":  5,
		"4":  4,
		"3":  3,
		"2":  2,
		"1":  1,
		"M":  0,
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, q *data.Queries) {
	var err error
	var rangeID int

	if rangeID, err = strconv.Atoi(r.FormValue("range-id")); err != nil {
		http.Error(w, "range-id not found", http.StatusBadRequest)
		return
	}

	arrowList := r.FormValue("arrow-list")

	cookie, err := r.Cookie("session_id")

	s, exists := auth.SessionMap[cookie.Value]
	if !exists {
		http.Error(w, "you are not authenticated", http.StatusBadRequest)
		return
	}

	// Remove white space then sepearate
	arrowList = strings.ReplaceAll(arrowList, " ", "")
	arrowList = strings.ReplaceAll(arrowList, "\n", "")
	arrowList = strings.ReplaceAll(arrowList, "\t", "")
	arrows := strings.Split(arrowList, ",")

	var totalScore int32 = 0
	for _, arrow := range arrows {
		if val, exists := scoresMap[arrow]; exists {
			totalScore += val
		} else {
			http.Error(w, "Invalid Arrows", http.StatusBadRequest)
			return
		}
	}

	// TODO: Add check if valid (right gender, age)

	// Start transactions so entire commit can be rolled back
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback() // auto rollback on return (after commit() doesn't do anything)

	qtx := q.WithTx(tx)

	res, err := qtx.CreateEnd(context.Background(), data.CreateEndParams{
		Rangeid:            int32(rangeID),
		Archeryaustraliaid: s.ArcheryAustraliaId,
		Finalscore:         int32(totalScore),
	})

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	for idx, arrow := range arrows {
		_, err := qtx.CreateScore(context.Background(), data.CreateScoreParams{
			Endid:       int32(id),
			Arrownumber: int32(idx + 1),
			Score:       arrow,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	tx.Commit()

	http.Redirect(w, r, "/", http.StatusFound)
}

func SubmitEventsHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {

	page := r.URL.Query().Get("page")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	events, err := q.GetAllEvents(context.Background(), data.GetAllEventsParams{
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
		Eventid:            int32(eventIDNum),
		Archeryaustraliaid: s.ArcheryAustraliaId,
		Limit:              10,
		Offset:             (int32(pageNum) - 1) * 10,
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

	ranges, err := q.GetRangeByRound(context.Background(), data.GetRangeByRoundParams{
		Roundid: int32(roundIDNum),
		Limit:   10,
		Offset:  (int32(pageNum) - 1) * 10,
	})

	type RangesWrapper struct {
		Rows    []data.Range
		Eventid string
		Page    int
	}

	eventID := r.URL.Query().Get("event_id")

	err = submitTmpl.ExecuteTemplate(w, "range-list", RangesWrapper{Rows: ranges, Page: pageNum, Eventid: eventID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	range_id := r.URL.Query().Get("range_id")
	if range_id == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err := submitTmpl.ExecuteTemplate(w, "submit-form", range_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
