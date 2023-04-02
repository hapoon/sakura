package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hapoon/sakura/handler"
)

func main() {
	mux := http.NewServeMux()

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	handler.Route(mux)

	fmt.Println("Access: http://localhost:8080/")
	if err := s.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
