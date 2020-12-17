package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data :=
		`{
  "message": "hello world", 
  "version": 0, 
  "request": {
    "host": "%s", 
    "url": "%s",
  }
}`
	fmt.Fprintf(w, data, r.Host, r.URL)
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Print("Starting helloworld REST API, version: 0")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
