package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/da-2115/ddp/web/data"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := os.Getenv("ARCHERY_DSN")
	if dsn == "" {
		dsn = "root:1234@tcp(127.0.0.1:3306)/ARCHERYDB?parseTime=true"
	}

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
	defer db.Close()

	slog.Info("Connected to DB")

	query := data.New(db)

	mux := http.NewServeMux()

	static := http.FileServer(http.Dir("static"))
	mux.Handle("GET /", static)

	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler(w, r, query)
	})

	mux.Handle("GET /api/login", authMiddleware(http.HandlerFunc(authTestHandler)))

	mux.Handle("GET /scores.html", authMiddleware(static))
	mux.Handle("GET /api/scores", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scoresHandler(w, r, query)
	})))

	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	slog.Info("Server listening", "address", "http://127.0.0.1"+srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
