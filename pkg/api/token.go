package api

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("my_secret_key")

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type claims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type tokenResponse struct {
	Token string `json:"token"`
	Username string `json:"username"`
	Role string `json:"role"`
}

type publicKeyResponse struct {
	PublicKey string `json:"publicKey"`
}

func (s *Server) publicKeyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&publicKeyResponse{PublicKey: string(jwtKey)})
}

func (s *Server) tokenGenerateHandler(w http.ResponseWriter, r *http.Request) {
	var credentials credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := credentials.Username
	role := "USER"

	if strings.Contains(username, "-") {
		split := strings.Split(username, "-")
		if len(split) > 0 {
			username = split[0]
			if (len(split[1])) > 0 {
				role = strings.ToUpper(split[1])
			}
		}
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &claims{
		Username: username,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		s.logger.Fatal("Token generate error", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tokenResponse{Token: tokenString, Username: username, Role: role})
}