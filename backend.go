package main

import (
	"fmt"
	"net/http"
)

func newBackendServer(port int) *http.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello from backend server on port %d!\n", port)
		if err != nil {
			return
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
}
