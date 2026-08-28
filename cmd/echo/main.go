package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgsamper/echo-service/internal/echo"
)

func main() {
	port, err := echo.ParsePort(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           http.HandlerFunc(echo.Handler),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("starting echo on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
