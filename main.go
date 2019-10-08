package main

import (
	"gorequests/pkg/http_session"
	"log"
)

func main() {
	http := http_session.New()
	r := http.EasyRequest("GET", "https://httpbin.org/ip")
	log.Printf("Response: %v", r.EasyJson())
	log.Printf("Response: %v", r.EasyString())
}
