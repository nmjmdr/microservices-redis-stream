package router

import (
	"account-service/src/datastore"
	"account-service/src/handlers/accounts"
	status "account-service/src/handlers/status"
	accountsService "account-service/src/services/accounts"
	authService "account-service/src/services/auth"
	"net/http"
	"os"

	authHandlers "account-service/src/handlers/auth"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Start starts the router and add routes
func Start(listenAddress string) {

	accountsDetailedStore, err1 := datastore.NewAccountsDetailedStore()
	accountStore, err2 := datastore.NewAccountStore()
	if err1 != nil {
		logrus.Errorf("Unable to connect to the data store, Error: %s", err1)
		os.Exit(1)
	}
	if err2 != nil {
		logrus.Errorf("Unable to connect to the data store, Error: %s", err2)
		os.Exit(1)
	}
	accountsSvc := accountsService.NewAccountsService(accountsDetailedStore, accountStore)
	accountsHandler := accounts.NewHandler(accountsSvc)

	authStore, err := datastore.NewAuthStore()
	if err != nil {
		logrus.Errorf("Unable to connect to the data store, Error: %s", err)
		os.Exit(1)
	}
	authSvc := authService.NewAuthService(authStore)
	authHandler := authHandlers.NewHandler(authSvc)

	r := mux.NewRouter()
	r.HandleFunc("/status", status.Handle).Methods("GET")
	r.HandleFunc("/accounts", AuthMiddleware(http.HandlerFunc(accountsHandler.Get))).Methods("GET")

	r.HandleFunc("/accounts/{accountID}/is-owned-by/{username}", accountsHandler.IsOwnedBy).Methods("GET")

	// auth routes
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	http.ListenAndServe(listenAddress, r)
}
