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
	"github.com/da-2115/ddp/web/middleware"
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

	// static files
	static := http.FileServer(http.Dir("static"))
	mux.Handle("GET /", static)

	// login api
	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginHandler(w, r, query)
	})
	mux.Handle("GET /api/login", auth.AuthMiddleware(http.HandlerFunc(auth.AuthTestHandler)))

	// View-Scores Page
	mux.Handle("GET /scores.html", auth.AuthMiddleware(static))
	mux.Handle("GET /components/scores", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.ScoresHandler(w, r, query)
	})))
	mux.HandleFunc("GET /components/nav", components.NavHandler)

	// Submit-Scores Page
	mux.Handle("GET /submit.html", auth.AuthMiddleware(static))
	mux.Handle("GET /components/events-list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.SubmitEventsHandler(w, r, query)
	})))
	mux.Handle("GET /components/rounds-list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.SubmitRoundsHandler(w, r, query)
	})))
	mux.Handle("GET /components/ranges-list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.SubmitRangesHandler(w, r, query)
	})))
	mux.Handle("GET /components/submit-form", auth.AuthMiddleware(http.HandlerFunc(components.SubmitFormHandler)))
	mux.Handle("POST /api/add-score", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.SubmitHandler(w, r, db, query)
	})))

	// View-All-Scores Page
	adminAuth := middleware.CreateStack(auth.AuthMiddleware, auth.AdminMiddleware) // checks auth then if admin
	mux.Handle("GET /view-all.html", adminAuth(static))
	mux.Handle("GET /components/view-all-events", adminAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.ViewAllEventsHandler(w, r, query)
	})))
	mux.Handle("GET /components/view-all-rounds", adminAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.ViewAllRoundsHandler(w, r, query)
	})))
	mux.Handle("GET /components/view-all-ranges", adminAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.ViewAllRangesHandler(w, r, query)
	})))
	mux.Handle("GET /components/view-all-ends", adminAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.ViewAllEndsHandler(w, r, query)
	})))
	mux.Handle("GET /components/view-all-scores", adminAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.ViewAllScoresHandler(w, r, query)
	})))

	// Define Server
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
