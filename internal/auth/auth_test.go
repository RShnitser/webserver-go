package auth

import (
	"testing"
	"github.com/google/uuid"
	"time"
)

func TestCheckPasswordHash(t *testing.T) {
	// First, we need to create some hashed passwords for testing
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJWTToken(t *testing.T){
	id := uuid.New()

	token, err := MakeJWT(id, "secret", time.Second * 100)
	if err != nil{
		t.Errorf("MakeJWT() error = %v", err)
	}

	wantID, err := ValidateJWT(token, "secret")
	if err != nil{
		t.Errorf("ValidateJWT() error = %v", err)
	}

	if id.String() != wantID.String(){
		t.Errorf("Validation failed")
	}
}

func TestGetBearerToken(t *testing.T){
	w := http.ResponseWriter{}
	w.Header().Set("Authorization", "Bearer token")
}