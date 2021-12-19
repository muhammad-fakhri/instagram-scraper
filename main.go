package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", PingHandler)
	mux.HandleFunc("/scrap", ScrapHandler)

	// start app
	log.Printf("starting app in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, RequestLogger(mux)))
}
