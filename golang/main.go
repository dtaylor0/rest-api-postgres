package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-api-postgres/internal/restdb"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "docker"
	SSLMODE  = "disable"
)

func main() {
	var connString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", HOST, USER, PASSWORD, DATABASE, SSLMODE)

	// connect to postgresql database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	// Test database conn
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful.")
	db.SetMaxOpenConns(10)

	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/api/msft", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Connection", "keep-alive")
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Add("Keep-Alive", "timeout=5")
		stmt := "select id, date, close, open, volume, high, low from hist_msft"
		res := restdb.Query(db, stmt)
		w.Write(res)
	})
	http.ListenAndServe(":8080", r)
}
