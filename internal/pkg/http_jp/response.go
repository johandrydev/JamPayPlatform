// Package http_jp provides HTTP response utilities for the JamPay application.
package http_jp

import (
	"encoding/json"
	"log"
	"net/http"

	"JamPay/internal/structure"
)

var emptyBody = []byte("{}")

// WriteError writes an error response with the specified status code and message.
func WriteError(w http.ResponseWriter, r *http.Request, status int, message string) {
	WriteJson(w, r, status, nil, message)
}

// WriteJson writes a JSON response with the specified status code, data, and message.
func WriteJson(w http.ResponseWriter, r *http.Request, status int, data any, message string) {
	response := structure.HttpResponse{
		Message: message,
	}
	if data != nil {
		response.Data = data
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding response in route: %s, err: %v", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(emptyBody)
		return
	}
}
