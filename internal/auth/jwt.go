package auth

import(
	"github.com/google/uuid"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
	"net/http"
	"strings"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: "chirpy",
		IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject: userID.String(),
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil{
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error){

	//claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return uuid.UUID{}, err
	} 
	
	claims, ok := token.Claims.(*jwt.RegisteredClaims);
	if !ok {
		return uuid.UUID{}, fmt.Errorf("Unknown claims type")
	}

	id, err := uuid.Parse(claims.Subject)
	if err != nil{
		return uuid.UUID{}, err
	}

	return id, nil
}

func GetBearerToken(headers http.Header) (string, error){
	authHeader := headers.Get("Authorization")
	if authHeader == ""{
		return "", fmt.Errorf("No Authorization header")
	}

	token := strings.Fields(authHeader)[1]
	return token, nil
}


