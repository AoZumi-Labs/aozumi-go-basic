package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(name string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Hello %s from the AOZUMI Go fixture!\n", name)
	})
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler(name)))
}
