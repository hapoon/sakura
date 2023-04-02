package handler

import "net/http"

func Route(mux *http.ServeMux) {
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello!\n"))
		})
	mux.HandleFunc("/echo", EchoHandler())
	mux.HandleFunc("/hello", HelloHandler())
}
