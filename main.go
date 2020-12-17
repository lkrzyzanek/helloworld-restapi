package main

import (
	"fmt"
	"log"
	"net/http"
)

var version = 0
var responseString = `{
  "message": "hello world", 
  "version": %v, 
  "request": {
    "host": "%s", 
    "url": "%s",
  }
}
`

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, responseString, version, r.Host, r.URL)
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Printf("Starting helloworld REST API, version: %v", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
