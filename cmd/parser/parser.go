package parser

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/gpraul/authorizer-challange/cmd/core"
)

// Parse function
func Parse(s string) (interface{}, error) {
	r := regexp.MustCompile(".+account.+")
	if r.MatchString(s) {
		account, _ := parseAccount(s)
		return account, nil
	}

	r = regexp.MustCompile(".+transaction.+")
	if r.MatchString(s) {
		transaction, _ := parseTransaction(s)
		return transaction, nil
	}

	return nil, errors.New("Formato não encontrado")
}

func parseAccount(s string) (core.Account, error) {
	a := &core.Account{}

	err := json.Unmarshal([]byte(s), a)
	if err != nil {
		return core.Account{}, errors.New("Erro ao converter para o tipo Account")
	}

	return *a, nil
}

func parseTransaction(s string) (core.Transaction, error) {
	t := &core.Transaction{}

	err := json.Unmarshal([]byte(s), t)
	if err != nil {
		return core.Transaction{}, errors.New("Erro ao converter para o tipo Transaction")
	}

	return *t, nil
}
