package accounts

import (
	"account-service/src/datastore"
	accountsDetailed "account-service/src/models/accountsdetailed"

	"github.com/pkg/errors"
)

// AccountsService - is used to get account details (TO DO: add methods to create new accounts)
type AccountsService interface {
	GetAll() ([]accountsDetailed.AccountDetailed, error)
	IsOwnedBy(userID string, accountID int) (bool, error)
}

type accountsSvc struct {
	accountsDetailedStore datastore.AccountsDetailedStore
	accountStore          datastore.AccountStore
}

// NewAccountsService - creates a new instance of AccountsService
func NewAccountsService(accountsDetailedStore datastore.AccountsDetailedStore,
	accountStore datastore.AccountStore,
) AccountsService {
	svc := &accountsSvc{
		accountsDetailedStore: accountsDetailedStore,
		accountStore:          accountStore,
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

func (a *accountsSvc) IsOwnedBy(userID string, accountID int) (bool, error) {
	results, err := a.accountStore.IsOwnedBy(userID, accountID)
	if err != nil {
		return false, errors.Wrap(err, "Unable to fetch account ownership details from database")
	}
	return results, nil
}
