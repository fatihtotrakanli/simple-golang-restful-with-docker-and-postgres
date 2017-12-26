package main

import (
	"net/http"
	"time"
	"Simple-GO-RestFul/muxes"
)

func main() {

	s := &http.Server{
		Addr:           ":3000",
		Handler:        muxes.SERVE(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
