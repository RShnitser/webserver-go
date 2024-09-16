package main

import (
	"fmt"
	"net/http"
)

func main(){
	port := "8080"
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr: ":" + port,
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("Starting server on port %s\n", port)
	err := server.ListenAndServe()
	if err != nil{
		fmt.Printf("Unable to start server: %v\n", err)
	}
}