package auth

// Ideally auth is best structured as a different service, but for now we are are including it
// as part of accounts service that runs over HTTPS
// all services have to run over HTTPS
import (
	"account-service/src/services/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// Ideally saved in secure store (s3) and injected into environment variable of servers
var jwtKey = []byte("a62f2225bf70bfaccbc7f1ef2a397836717377de")

// Handler - accounts handler
type Handler struct {
	svc auth.AuthService
}

// Claims - jwt claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewHandler - returns a new instance for handler
func NewHandler(svc auth.AuthService) *Handler {
	return &Handler{
		svc: svc,
	}
}

// Credentials - request payload for login
type Credentials struct {
	Username string
	Password string
}

func generateJWT(username string) (string, error) {
	// move it to config later
	// ideally set a low expiration period (5 min) and period a refresh method
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

// Login - login using the credentials
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := Credentials{}
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok, err := h.svc.Login(credentials.Username, credentials.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to login, server error")
		logrus.Error("Unable to login, server error", err)
		return
	}

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := generateJWT(credentials.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to generate JWT token, server error")
		logrus.Error("Unable to generate JWT token, server error", err)
		return
	}
	// return the token, the token has to be set as Authorization Bearer token in API calls
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, tokenString)
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	// TO DO: complete this method later
}
