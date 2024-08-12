package main

import (
	"log"
	"net/http"
	"time"

	"github.com/drizzleent/wallet-tracker/backend/internal/api/handler"
)

const (
	adrr = "0x5295AFCE96E05C716d3C415236572DBAB9b5dA92"
)

func main() {

	srv := &http.Server{
		Addr:           "localhost:8001",
		Handler:        handler.NewHandler().InitRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("faile to run server %s", err.Error())
	}
}
