package http_transport

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, response Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
