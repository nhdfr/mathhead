package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type TwoNumReq struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var req TwoNumReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode request", "errorr", err)
		RespondWithWriter(w, http.StatusBadRequest, "Invalid reques body")
		return
	}

}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("Failed to encode response", "Error", err)
	}
}

func RespondWithWriter(w http.ResponseWriter, statusCode int, message string) {
	respondWithJson(w, statusCode, ErrorResponse{Error: message})
}
