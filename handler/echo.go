package handler

import (
	"fmt"
	"log"
	"net/http"
)

func EchoHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("ContentLength:", r.ContentLength)
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		log.Println("body:", string(body))
		fmt.Fprintln(w, string(body))
	}
}
