package auth

import (
	datastore "account-service/src/datastore"
	"crypto/sha1"
	"fmt"

	"github.com/pkg/errors"
)

type AuthService interface {
	Login(username string, password string) (bool, error)
}

// Uses SHA-1 hash
// Ideally should used salt too
type authService struct {
	authStore datastore.AuthStore
}

func NewAuthService(authStore datastore.AuthStore) AuthService {
	return &authService{
		authStore: authStore,
	}
}

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (a *authService) Login(username string, password string) (bool, error) {
	hashed := hash(password)
	ok, err := a.authStore.Match(username, hashed)
	if err != nil {
		return false, errors.Wrap(err, "Unable to match username/password against data in DB")
	}
	return ok, nil
}
