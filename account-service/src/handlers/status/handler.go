package status

import (
	"fmt"
	"net/http"
)

// Handle - health check end point
// Can be enhanced to collect and report status of various parameters
func Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
