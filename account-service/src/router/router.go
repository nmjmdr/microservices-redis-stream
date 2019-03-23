package router

import (
	status "account-service/src/handlers/status"
	"net/http"

	"github.com/gorilla/mux"
)

// Start starts the router and add routes
func Start(listenAddress string) {
	r := mux.NewRouter()

	r.HandleFunc("/status", status.Handle).Methods("GET")

	http.ListenAndServe(listenAddress, r)
}
