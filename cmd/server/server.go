package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/hsmtkk/heroku-go-web-app/pkg/webapp"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/post", webapp.New().GenerateHandler())
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
