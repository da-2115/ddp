package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"

	"github.com/da-2115/ddp/web/data"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/ARCHERYDB?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	slog.Info("Connected to DB")

	query := data.New(db)

	mux := http.NewServeMux()

	static := http.FileServer(http.Dir("static"))
	mux.Handle("GET /", static)

	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler(w, r, query)
	})

	mux.Handle("GET /api/login", authMiddleware(http.HandlerFunc(authTestHandler)))

	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	slog.Info("Server listening", "address", "http://127.0.0.1" + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	// pswd, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = query.DeleteMember(context.Background(), "1234")
	// if err != nil {
	// 	slog.Error("Failed to DeleteMember", "err", err)
	// }
	//
	//
	// dob, _ := time.Parse("2006-01-02", "2002-02-26")
	// _, err = query.CreateMember(context.Background(), data.CreateMemberParams{
	// 	Archeryaustraliaid: "1234",
	// 	Passwordhash: string(pswd),
	// 	Firstname: "Luke",
	// 	Dateofbirth: dob,
	// 	Gender: true,
	// 	Clubrecorder: false,
	// 	Defaultdivision: "Compound",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// res, err := query.GetMemberByID(context.Background(), "1234")
	// if err != nil {
	// 	log.Fatal("Failed to read ", err)
	// }
	// fmt.Println(res)
}
