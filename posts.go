package main

import (
	"net/http"
	"encoding/json"
	"strings"
)

func replaceProfane(input string)string{
	badWords := map[string]struct{}{
		"kerfuffle":{},
		"sharbert":{},
		"fornax":{},
	}

	words := strings.Split(input, " ")
	for i, word := range words{
		_, ok := badWords[strings.ToLower(word)]
		if  ok{
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}

func handleValidateChirps(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respBody := returnVals{
		CleanedBody: replaceProfane(params.Body),
	}
	respondWithJSON(w, http.StatusOK, respBody)
}
