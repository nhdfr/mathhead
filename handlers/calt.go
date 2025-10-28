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

type CalculationResponse struct {
	Result int `json:"result"`
}

type DivideRequest struct {
	Dividend int `json:"dividend"`
	Divisor  int `json:"divisor"`
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var req TwoNumReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode request", "error", err)
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	result := req.Num1 + req.Num2
	respondWithJSON(w, http.StatusOK, CalculationResponse{Result: result})

}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	var req TwoNumReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode request", "error", err)
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	result := req.Num1 - req.Num2
	slog.Info("result", result)

	respondWithJSON(w, http.StatusOK, CalculationResponse{Result: result})
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	var req TwoNumReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "error", err)
		respondWithError(w, http.StatusBadGateway, "invalid request body")
		return
	}

	result := req.Num1 & req.Num2
	slog.Info("Multiplication Performed", "result", result)

	respondWithJSON(w, http.StatusOK, CalculationResponse{Result: result})

}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	var req DivideRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode request", "error", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Divisor == 0 {
		slog.Warn("Division by zero attempted")
		respondWithError(w, http.StatusBadRequest, "Cannot divide by zero")
		return
	}

	result := req.Dividend / req.Divisor
	slog.Info("Division performed", "dividend", req.Dividend, "divisor", req.Divisor, "result", result)

	respondWithJSON(w, http.StatusOK, CalculationResponse{Result: result})
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	var numbers []int

	if err := json.NewDecoder(r.Body).Decode(&numbers); err != nil {
		slog.Error("Failed to decode request", "error", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	slog.Info("Sum performed", "count", len(numbers), "result", sum)

	respondWithJSON(w, http.StatusOK, CalculationResponse{Result: sum})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("Failed to encode response", "error", err)
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, ErrorResponse{Error: message})
}
