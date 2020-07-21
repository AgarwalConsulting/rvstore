package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const jwtSigningKey = "s8&5rtrjGSrgY$6U2UcvU8qbYtz8%Qd7Y*g8ar9Qk^h&EmJk7fj$R&BSyDD4bQmjP73zF5#F6Pf^6!F@qP5BdpBJEmBAnJD3aRs"

// AccessToken is a JWT sent back to the user
type AccessToken struct {
	Token string `json:"access_token"`
}

// GenerateJwt creates a new signed JWT
func GenerateJwt() AccessToken {
	expires := time.Now().Unix()
	expires += 60 * 60 * int64(12) // Expire in 12 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       "api.rvstore.com",
		"exp":       expires,
		"sub":       "12345", // This is the user ID
		"firstName": "John",
		"lastName":  "Doe",
		"username":  "algogrit",
	})

	// Sign and get the complete encoded token as a string using the secret
	// This is a shared secret that your services should know
	tokenString, err := token.SignedString([]byte(jwtSigningKey))

	log.Info(tokenString, err)

	accessToken := AccessToken{
		Token: tokenString,
	}

	return accessToken
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	token := GenerateJwt()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(token)
}

func main() {
	httpAddr := ":9003"
	r := mux.NewRouter()

	r.HandleFunc("/auth/login", loginHandler)

	r.Use(handlers.CORS())

	log.Infof("Starting auth service on: %s...\n", httpAddr)
	http.ListenAndServe(httpAddr, handlers.LoggingHandler(os.Stdout, r))
}
