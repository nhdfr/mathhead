package main

import (
	"calculator/handlers"
	"log/slog"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /add", handlers.AddHandler)
	mux.HandleFunc("POST /subtract", handlers.SubtractHandler)
	mux.HandleFunc("POST /multiply", handlers.MultiplyHandler)
	mux.HandleFunc("POST /divide", handlers.DivideHandler)
	mux.HandleFunc("POST /sum", handlers.SumHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	port := ":3000"
	slog.Info("Server starting", "port", port)

	if err := http.ListenAndServe(port, handler); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
