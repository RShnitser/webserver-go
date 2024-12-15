package main

import(
	"net/http"
	"server/internal/auth"
	"time"
)

func(cfg *apiConfig) handleRefresh(w http.ResponseWriter, r *http.Request) {

	type returnVals struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized, "Missing header", err)
		return
	}

	refreshToken, err := cfg.db.GetRefreshToken(r.Context(), token)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized, "Invalid Token", err)
		return
	}

	if time.Now().UTC().After(refreshToken.ExpiresAt){
		respondWithError(w, http.StatusUnauthorized, "Token is expired", err)
		return
	}

	jwtToken, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Could not create token", err)
		return
	}

	respBody := returnVals{
		Token: jwtToken,
	}
	respondWithJSON(w, http.StatusOK, respBody)
}