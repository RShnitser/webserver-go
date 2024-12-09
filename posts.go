package main

import(
	"fmt"
	"net/http"
)

func handleValidatePosts(w http.ResponseWriter, r *http.Request){
	type parameters struct {
        Body string `json:"body"`
    }

	decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
       
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error decoding parameters: %s", err)))
    }
}