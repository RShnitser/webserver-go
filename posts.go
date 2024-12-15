package main

import (
	"net/http"
	"encoding/json"
	"strings"
	"github.com/google/uuid"
	"time"
	"server/internal/database"
	"server/internal/auth"
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

func(cfg *apiConfig) handleAddChip(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	token, err := auth.GetBearerToken()

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

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{params.Body, params.UserID})
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
