package customers

import (
	"customer-service/src/handlers"
	customerModel "customer-service/src/models/customers"
	"encoding/json"
	"fmt"
	"net/http"

	"customer-service/src/datastore"
	"customer-service/src/messaging"
	"customer-service/src/services"

	"github.com/sirupsen/logrus"
)

// Handler - defines operations on Customer resource
type Handler struct {
	publisher messaging.Publisher
	dataStore datastore.CustomerStore
}

// NewCustomerHandler - returns a new handler instance for Customer resource
func NewCustomerHandler(publisher messaging.Publisher, dataStore datastore.CustomerStore) *Handler {
	return &Handler{
		publisher: publisher,
		dataStore: dataStore,
	}
}

// Create - create a new customer
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	data, err := handlers.ReadRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to read request")
		logrus.Error("Unable to read request during Create Customer", err)
		return
	}

	customer, err := customerModel.Deserialize(data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to derserialize customer from request stream")
		return
	}

	err = customerModel.ValidateInput(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("Invalid data: %s", err))
		return
	}

	svc := services.NewCustomerService(h.publisher, h.dataStore)
	id, err := svc.Create(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to create customer, internal error")
		logrus.Error("Unable to create customer, internal error", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	result := handlers.CreateResult{ID: id}
	bytes, _ := json.Marshal(result)
	fmt.Fprintf(w, string(bytes))
}
