package status

import (
	"fmt"
	"net/http"
)

// Handle - health check end point
// Can be enhanced to collect and report status of various parameters
func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
