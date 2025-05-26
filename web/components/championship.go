package components

import (
	"context"
	"net/http"
	"strconv"

	"github.com/da-2115/ddp/web/data"
)

func ViewChampionshipEventsHandler(w http.ResponseWriter, r *http.Request, q *data.Queries) {
	var page int
	var err error
	if page, err = strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	type wrapper struct {
		Rows []data.GetChampionshipEventsRow
		Page int
	}

	events, err := q.GetChampionshipEvents(context.Background(), data.GetChampionshipEventsParams{
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
