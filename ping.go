package main

import "net/http"

func PingHandler(w http.ResponseWriter, req *http.Request) {
	SendJsonResponse(w, "Pong", nil)
}
