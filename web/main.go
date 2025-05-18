package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/da-2115/ddp/web/auth"
	"github.com/da-2115/ddp/web/components"
	"github.com/da-2115/ddp/web/data"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Gets the database connection string
	dsn := os.Getenv("ARCHERY_DSN")
	if dsn == "" {
		dsn = "root:1234@tcp(127.0.0.1:3306)/ARCHERYDB?parseTime=true"
	}

	// Connect to the database / make sure it is reachable
	var db *sql.DB
	var err error
	for {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}
		slog.Info("Waiting for database..", "err", err)
		time.Sleep(2 * time.Second)
	}
	defer db.Close() // defer just means when main() ends run this func

	slog.Info("Connected to DB")

	// Setting up db with on SQLC library
	query := data.New(db)

	// set up mux eg. the http requests and what funcs they should call
	mux := http.NewServeMux()

	static := http.FileServer(http.Dir("static"))
	mux.Handle("GET /", static)

	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginHandler(w, r, query)
	})

	mux.Handle("GET /api/login", auth.AuthMiddleware(http.HandlerFunc(auth.AuthTestHandler)))

	mux.Handle("GET /scores.html", auth.AuthMiddleware(static))
	mux.Handle("GET /components/scores", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.ScoresHandler(w, r, query)
	})))
	mux.HandleFunc("GET /components/nav", components.NavHandler)

	mux.Handle("GET /submit.html", auth.AuthMiddleware(static))
	mux.Handle("GET /components/events-list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			components.SubmitEventsHandler(w,r,query)
		})))
	mux.Handle("GET /components/rounds-list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			components.SubmitRoundsHandler(w,r,query)
		})))
	mux.Handle("GET /components/ranges-list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			components.SubmitRangesHandler(w,r,query)
		})))

	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	// runs the server
	slog.Info("Server listening", "address", "http://127.0.0.1"+srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
