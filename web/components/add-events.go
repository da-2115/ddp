package components

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/da-2115/ddp/web/data"
	"github.com/da-2115/ddp/web/util"
)

var addEventTmpl *template.Template

func init() {
	addEventTmpl = util.Unwrap(template.ParseFiles("views/add-event/form.html"))
}

type eventRange struct {
	targetSize int
	distance   int
}

type eventRound struct {
	ranges   []eventRange
	class    string
	division string
	gender   string
}

func AddEventHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, q *data.Queries) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	var eventName string
	eventName = r.FormValue("event-name")
	if eventName == "" {
		http.Error(w, "Event name empty", http.StatusBadRequest)
		return
	}

	var eventDate time.Time
	eventDate, err = time.Parse("2006-01-02", r.FormValue("event-date"))
	if err != nil {
		http.Error(w, "Date is not valid", http.StatusBadRequest)
		return
	}

	// Start transactions so entire commit can be rolled back
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback() // auto rollback on return (after commit() doesn't do anything)

	qtx := q.WithTx(tx)

	res, err := qtx.CreateEvent(context.Background(), data.CreateEventParams{
		Name: eventName,
		Date: eventDate,
	})
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusBadRequest)
		return
	}

	EventID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var eventRounds []eventRound

	var done []string // list of already completed rounds

	// goes through the form and collects rounds and ranges
	for name := range r.Form {
		if name[0:6] == "round-" {

			valLen := 1
			found := false
			for _, r := range name[7:] {
				if r != '-' {
					valLen++
				} else {
					found = true
					break
				}
			}

			contains := false
			for _, c := range done {
				if c == name[6:6+valLen] {
					contains = true
				}
			}

			if contains {
				continue
			} else {
				done = append(done, name[6:6+valLen])
			}

			if !found {
				http.Error(w, "", http.StatusBadRequest)
				return
			}

			class := r.FormValue("round-" + name[6:6+valLen] + "-class")
			division := r.FormValue("round-" + name[6:6+valLen] + "-division")
			gender := r.FormValue("round-" + name[6:6+valLen] + "-gender")

			rangeTargets := r.Form["range-"+name[6:6+valLen]+"-distance[]"]
			rangeDistances := r.Form["range-"+name[6:6+valLen]+"-target[]"]

			round := eventRound{
				class:    class,
				division: division,
				gender:   gender,
			}

			for idx := range rangeTargets {
				var target int
				var distance int
				var err error
				if target, err = strconv.Atoi(rangeTargets[idx]); err != nil {
					http.Error(w, "", http.StatusBadRequest)
					return
				}
				if distance, err = strconv.Atoi(rangeDistances[idx]); err != nil {
					http.Error(w, "", http.StatusBadRequest)
					return
				}

				r := eventRange{
					targetSize: target,
					distance:   distance,
				}

				round.ranges = append(round.ranges, r)
			}

			eventRounds = append(eventRounds, round)

			continue
		}
	}

	for _, round := range eventRounds {
		res, err := qtx.CreateRound(context.Background(), data.CreateRoundParams{
			Gender:   data.RoundGender(round.gender),
			Class:    data.RoundClass(round.class),
			Division: data.RoundDivision(round.division),
			Eventid:  int32(EventID),
		})
		if err != nil {
			http.Error(w, "Failed to create round", http.StatusBadRequest)
			return
		}

		roundID, err := res.LastInsertId()

		for _, roundRange := range round.ranges {
			_, err := qtx.CreateRange(context.Background(), data.CreateRangeParams{
				Roundid:    int32(roundID),
				Targetsize: int32(roundRange.targetSize),
				Distance:   int32(roundRange.distance),
			})
			if err != nil {
				http.Error(w, "Failed to create round", http.StatusBadRequest)
				return
			}
		}
	}

	tx.Commit()
}

func AddEventFormHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	addEventTmpl.ExecuteTemplate(w, "form", 0)
}

func AddRoundFormHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var roundNum int
	var err error
	if roundNum, err = strconv.Atoi(r.URL.Query().Get("num")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	addEventTmpl.ExecuteTemplate(w, "round", roundNum)
}

func AddRangeFormHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var rangeNum int
	var err error
	if rangeNum, err = strconv.Atoi(r.URL.Query().Get("num")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	addEventTmpl.ExecuteTemplate(w, "range", rangeNum)
}
