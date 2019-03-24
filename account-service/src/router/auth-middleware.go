package router

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// AuthMiddleware function, which will be called for each request
func AuthMiddleware(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			http.Error(w, "Missing Bearer token in request", http.StatusBadRequest)
			return
		}
		parts := strings.Split(token, " ")
		if len(parts) != 2 {
			http.Error(w, "Missing Bearer token in request", http.StatusBadRequest)
			return
		}
		result, err := VerifyToken(parts[1])

		if err != nil {
			logrus.Errorf("Error while trying to verify token: %s", err)
			http.Error(w, "Internal server error, while trying to verify token", http.StatusInternalServerError)
			return
		}

		if result.OK {
			// set the user name on the request header and pass the request forward
			r.Header.Set("username", result.Username)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	}
}
