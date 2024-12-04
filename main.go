package main

import (
	"fmt"
	"net/http"
)

func handleReadiness(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", " text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main(){
	filePathRoot := "."
	port := "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot))))
	mux.HandleFunc("/healthz", handleReadiness)
	
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