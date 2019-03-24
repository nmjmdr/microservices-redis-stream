package router

import (
	"customer-service/src/datastore"
	customersHandler "customer-service/src/handlers/customers"
	invoicesHandler "customer-service/src/handlers/invoices"
	"customer-service/src/handlers/status"
	"customer-service/src/messaging"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Start starts the router and add routes
func Start(listenAddress string) {
	r := mux.NewRouter()

	customerDatastore, customerStoreErr := datastore.NewCustomerStore()
	invoiceDatastore, invoiceStoreErr := datastore.NewInvoiceStore()
	if customerStoreErr != nil || invoiceStoreErr != nil {
		logrus.Error("Unable to connect to customer database")
		os.Exit(1)
	}

	customerHandler := customersHandler.NewCustomerHandler(messaging.NewRedisStreamPublisher(), customerDatastore)
	invoiceHandler := invoicesHandler.NewInvoiceHandler(messaging.NewRedisStreamPublisher(), invoiceDatastore)

	r.HandleFunc("/status", status.Handle).Methods("GET")
	r.HandleFunc("/customers", AuthMiddleware(http.HandlerFunc(customerHandler.Create))).Methods("POST")
	r.HandleFunc("/invoices", AuthMiddleware(http.HandlerFunc(invoiceHandler.Create))).Methods("POST")

	http.ListenAndServe(listenAddress, r)
}
