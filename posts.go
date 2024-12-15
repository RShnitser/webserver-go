package main

import (
	"net/http"
	"encoding/json"
	"strings"
	"github.com/google/uuid"
	"time"
	"server/internal/database"
	"server/internal/auth"
	"errors"
)

type Chirp struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body string `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

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

func validateChirp(body string)(string, error) {
	
	if len(body) > 140 {
		return "", errors.New("Chirp is too long")
	}

	result := replaceProfane(body)
	return result, nil
}

func(cfg *apiConfig) handleAddChip(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized, "Missing header", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized, "Invalid Token", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	validChirp, err := validateChirp(params.Body)
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "Invalid chirp", err)
		return
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{validChirp, userID})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
		return
	}

	respBody := Chirp{
		ID: chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body: chirp.Body,
		UserID: chirp.UserID,
	}
	respondWithJSON(w, http.StatusCreated, respBody)
}

func(cfg *apiConfig) handleGetAllChirps(w http.ResponseWriter, r *http.Request) {
	
	chirps, err := cfg.db.GetChrips(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}

	respBody := []Chirp{}
	for _, chirp := range chirps{
		respBody = append(respBody, Chirp{
			ID: chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body: chirp.Body,
			UserID: chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, respBody)

}

func(cfg *apiConfig) handleGetChirpByID(w http.ResponseWriter, r *http.Request) {
	
	idString := r.PathValue("chirpID")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse id", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}

	respBody := Chirp{
		ID: chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body: chirp.Body,
		UserID: chirp.UserID,
	}
	respondWithJSON(w, http.StatusOK, respBody)

}
