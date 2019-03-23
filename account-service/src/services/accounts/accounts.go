package accounts

import (
	"account-service/src/datastore"
	accountsDetailed "account-service/src/models/accountsdetailed"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// AccountsService - is used to get account details (TO DO: add methods to create new accounts)
type AccountsService interface {
	GetAll() ([]accountsDetailed.AccountDetailed, error)
}

type accountsSvc struct {
	accountsDetailedStore datastore.AccountsDetailedStore
}

// NewAccountsService - creates a new instance of AccountsService
func NewAccountsService(accountsDetailedStore datastore.AccountsDetailedStore) AccountsService {
	svc := &customerSvc{
		accountsDetailedStore: accountsDetailedStore,
	}
	return svc
}

func (a *AccountsService) GetAll() ([]accountsDetailed.AccountDetailed, error) {
	return []accountsDetailed.AccountDetailed, nil
}