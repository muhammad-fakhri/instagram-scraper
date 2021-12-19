package main

import "net/http"

func PingHandler(w http.ResponseWriter, req *http.Request) {
	SendSuccessResponse(w, "Pong", nil)
}
