package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handleWriteHits(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", " text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
}

func main(){
	filePathRoot := "."
	port := "8080"

	cfg := apiConfig{}

	mux := http.NewServeMux()
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot)))))
	mux.HandleFunc("/healthz", handleReadiness)
	mux.HandleFunc("/metrics", cfg.handleWriteHits)
	mux.HandleFunc("/reset", cfg.handleReset)
	
	server := http.Server{
		Handler: mux,
		Addr: ":" + port,
	}
	fmt.Printf("Starting server on port %s\n", port)
	err := server.ListenAndServe()
	if err != nil{
		fmt.Printf("Unable to start server: %v\n", err)
	}
}