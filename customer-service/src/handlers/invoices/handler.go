package customers

import (
	"customer-service/src/handlers"
	invoiceModel "customer-service/src/models/invoices"
	"encoding/json"
	"fmt"
	"net/http"

	"customer-service/src/datastore"
	"customer-service/src/messaging"
	"customer-service/src/services"

	"github.com/sirupsen/logrus"
)

// Handler - invoice operations
type Handler struct {
	publisher messaging.Publisher
	dataStore datastore.InvoiceStore
}

// NewInvoiceHandler - returns a new handler instance for invoice resource
func NewInvoiceHandler(publisher messaging.Publisher, dataStore datastore.InvoiceStore) *Handler {
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
		logrus.Error("Unable to read request during Create Invoice", err)
		return
	}

	invoice, err := invoiceModel.Deserialize(data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to derserialize invoice from request stream %s", err)
		return
	}

	err = invoiceModel.ValidateInput(invoice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("Invalid data: %s", err))
		return
	}

	svc := services.NewInvoiceService(h.publisher, h.dataStore)
	id, err := svc.Create(invoice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to create Invoice, internal error")
		logrus.Error("Unable to create Invoice, internal error", err)
		return
	}
	result := handlers.CreateResult{ID: id}
	bytes, _ := json.Marshal(result)
	fmt.Fprintf(w, string(bytes))
}
