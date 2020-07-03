package adapter

import (
	"errors"

	"github.com/gpraul/authorizer-challange/cmd/core"
)

var accounts []core.Account = []core.Account{}

// FindOneAccount function
func FindOneAccount() (core.Account, error) {
	lenght := len(accounts)
	if lenght > 0 {
		return accounts[0], nil
	}

	return core.Account{}, errors.New("Nenhuma conta encontrada")
}

// SaveAccount function
func SaveAccount(a core.Account) {
	accounts = append(accounts, a)
}
