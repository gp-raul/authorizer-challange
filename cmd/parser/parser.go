package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/gpraul/authorizer-challange/cmd/core"
)

func Parse(s string) {
	r := regexp.MustCompile(".+account.+")
	if r.MatchString(s) {
		account, _ := parseAccount(s)
		fmt.Println(account)
		return
	}

	r = regexp.MustCompile(".+transaction.+")
	if r.MatchString(s) {
		transaction, _ := parseTransaction(s)
		fmt.Println(transaction)
		return
	}

	fmt.Println("Formato não encontrado...")
}

func parseAccount(s string) (*core.Account, error) {
	a := &core.Account{}

	err := json.Unmarshal([]byte(s), a)
	if err != nil {
		return nil, errors.New("Erro ao converter para o tipo Account")
	}

	return a, nil
}

func parseTransaction(s string) (*core.Transaction, error) {
	t := &core.Transaction{}

	err := json.Unmarshal([]byte(s), t)
	if err != nil {
		return nil, errors.New("Erro ao converter para o tipo Transaction")
	}

	return t, nil
}
