package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	filePathRoot := "."
	port := "8080"

	cfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	fsHandler := cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot))))

	mux := http.NewServeMux()
	mux.Handle("/app/", fsHandler)
	mux.HandleFunc("GET /healthz", handleReadiness)
	mux.HandleFunc("GET /admin/metrics", cfg.handleWriteHits)
	mux.HandleFunc("POST /admin/reset", cfg.handleReset)

	mux.HandleFunc("POST /api/validate_chirp", handleValidateChirps)

	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	fmt.Printf("Starting server on port %s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Unable to start server: %v\n", err)
	}
}
