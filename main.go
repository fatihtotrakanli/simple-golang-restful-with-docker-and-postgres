package main

import (
	"net/http"
	"time"
	"Simple-GO-RestFul/muxes"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"Simple-GO-RestFul/config"
)

func main() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.PORT)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	log.Printf("Postgres started at %s PORT", config.PORT)
	defer db.Close()

	s := &http.Server{
		Addr:           ":3000",
		Handler:        muxes.SERVE(db),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}