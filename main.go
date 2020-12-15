package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = 2.0
var responseString = `{
  "message": "hello world", 
  "version": %v, 
  "pod": %s,
  "request": {
    "host": "%s", 
    "url": "%s",
  }
}
`

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pod, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, responseString, version, pod, r.Host, r.URL)
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Printf("Starting helloworld REST API, version: %v", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
