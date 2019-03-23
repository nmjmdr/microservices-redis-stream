package accounts

import (
	"fmt"
	"net/http"
)

// Get - get accounts
func Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
