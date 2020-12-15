package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world - LATEST", "version": 0}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Print("Starting helloworld REST API, version: 0")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
