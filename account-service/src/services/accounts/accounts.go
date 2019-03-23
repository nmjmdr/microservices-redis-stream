package accounts

import (
	"account-service/src/datastore"
	accountsDetailed "account-service/src/models/accountsdetailed"

	"github.com/pkg/errors"
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
	svc := &accountsSvc{
		accountsDetailedStore: accountsDetailedStore,
	}
	return svc
}

func (a *accountsSvc) GetAll() ([]accountsDetailed.AccountDetailed, error) {
	results, err := a.accountsDetailedStore.GetAll()
	if err != nil {
		return []accountsDetailed.AccountDetailed{}, errors.Wrap(err, "Unable to fetch account detail results from database")
	}
	return results, nil
}
