package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func LogError(w http.ResponseWriter, err error) {
	fmt.Println("[Error Log] " + err.Error())
	SendErrorResponse(w, err)
}

func sendResponse(w http.ResponseWriter, message string, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(JsonHTTPResponse{
		Message: message,
		Data:    data,
	})
}

func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	sendResponse(w, message, data, http.StatusOK)
}

func SendErrorResponse(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	switch err {
	case ErrReqQueryInvalid:
		code = http.StatusBadRequest
	}

	sendResponse(w, err.Error(), nil, code)
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}
