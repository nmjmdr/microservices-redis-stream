package accounts

import (
	"account-service/src/services/accounts"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	results, err := h.svc.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to get all account details")
		logrus.Error("Unable to get all account details", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(results)
	fmt.Fprintf(w, string(bytes))
}

// Result - result of is owned by. Ideally all these are abstracte into a response package
type Result struct {
	OK bool `json:"ok"`
}

// IsOwnedBy - checks if an account is owned by a given user
func (h *Handler) IsOwnedBy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIDString, _ := vars["accountID"]
	username, _ := vars["username"]
	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid account id passed as parameter")
		return
	}
	ok, err := h.svc.IsOwnedBy(username, accountID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to check ownership of account")
		logrus.Error("Unable to chech ownership of account in IsOwnedBy method", err)
		return
	}
	result := Result{
		OK: ok,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(result)
	fmt.Fprintf(w, string(bytes))
}
