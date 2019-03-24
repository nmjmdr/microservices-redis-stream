package accounts

import (
	"account-service/src/services/accounts"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Handler - accounts handler
type Handler struct {
	svc accounts.AccountsService
}

// NewHandler - returns a new instance for handler
func NewHandler(svc accounts.AccountsService) *Handler {
	return &Handler{
		svc: svc,
	}
}

// Get - get accounts
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	results, err := h.svc.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to get all account details")
		logrus.Error("Unable to get all account details", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(results)
	fmt.Fprintf(w, string(bytes))
}

func (h *Handler) IsOwnbedBy(w http.ResponseWriter, r *http.Request) {

}
