package main

import _ "github.com/lib/pq"
import (
	"fmt"
	"net/http"
	"sync/atomic"
	"server/internal/database"
	"database/sql"
	"os"
	
	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db *database.Queries
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL not found")
		return
	}
	dbConnection, err := sql.Open("postgres", dbURL)
	if dbURL != nil {
		fmt.Printfln("Could not connect to database: %s", err)
		return
	}

	filePathRoot := "."
	port := "8080"

	cfg := apiConfig{
		fileserverHits: atomic.Int32{},
		dbQueries : database.New(dbConnection),
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
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Unable to start server: %v\n", err)
	}
}
