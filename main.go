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
	platform string
	jwtSecret string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL not found")
		return
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		fmt.Println("PLATFORM must be set")
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET must be set")
		return
	}

	dbConnection, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Could not connect to database: %s\n", err)
		return
	}

	filePathRoot := "."
	port := "8080"

	cfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db : database.New(dbConnection),
		platform: platform,
		jwtSecret: jwtSecret,
	}

	fsHandler := cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot))))

	mux := http.NewServeMux()
	mux.Handle("/app/", fsHandler)
	mux.HandleFunc("GET /healthz", handleReadiness)
	mux.HandleFunc("GET /admin/metrics", cfg.handleWriteHits)
	mux.HandleFunc("POST /admin/reset", cfg.handleReset)

	mux.HandleFunc("POST /api/chirps", cfg.handleAddChip)
	mux.HandleFunc("GET /api/chirps", cfg.handleGetAllChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handleGetChirpByID)

	mux.HandleFunc("POST /api/users", cfg.handleAddUser)

	mux.HandleFunc("POST /api/login", cfg.handleLogin)

	mux.HandleFunc("POST /api/refresh", cfg.handleRefresh)


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
