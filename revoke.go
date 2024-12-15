package main

import(
	"net/http"
	"server/internal/auth"
)

func(cfg *apiConfig) handleRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized, "Missing header", err)
		return
	}

	_, err = cfg.db.RevokeRefreshToken(r.Context(), token)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Invalid Token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
