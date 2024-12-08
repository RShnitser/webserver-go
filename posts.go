package main

import(
	"net/http"
)

func handlePosts(w http.ResponseWriter, r *http.Request){
	type chirp struct {
        Body string `json:"body"`
    }
}