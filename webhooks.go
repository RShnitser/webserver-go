package main

import (
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"database/sql"
	"errors"
	"server/internal/auth"
)

func (cfg *apiConfig) handleUpgrade(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	key, err := auth.GetAPIKey(r.Header)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized, "Missing header", err)
		return
	}

	if cfg.polkaKey != key{
		respondWithError(w, http.StatusUnauthorized, "invalid polka key", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded"{
		w.WriteHeader(http.StatusNoContent)
		return
	}

	id, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse id", err)
		return
	}

	_, err = cfg.db.UpdateChirpyRed(r.Context(), id)
	if err != nil{
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Could not find user", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not update user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}