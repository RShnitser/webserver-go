package main

import(
	"net/http"
	"server/internal/auth"
	"server/internal/database"
	"encoding/json"
	"time"
	"github.com/google/uuid"
)

func(cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	type returnVals struct {
		ID uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email string `json:"email"`
		Token string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Could not create token", err)
		return
	}

	refreshTokenString, err := auth.MakeRefreshToken()
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Could not create refresh token", err)
		return
	}

	refreshToken, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{refreshTokenString, user.ID, time.Now().UTC().Add(60 * 24 * time.Hour)})
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Could not create refresh token", err)
		return
	}

	respBody := returnVals{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
		Token: token,
		RefreshToken: refreshToken.Token,
	}
	respondWithJSON(w, http.StatusOK, respBody)
}