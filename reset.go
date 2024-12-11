package main

import (
	"net/http"
)

func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev"{
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment."))
		return
	}
	cfg.fileserverHits.Store(0) 

	err := cfg.db.DeleteUsers(r.Context())
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete user", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reset"))
}
