package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/gpraul/authorizer-challange/internal/app/account"
)

func main() {
	// recebendo um json
	jsonAccount := `{"account": {"active-card": true, "available-limit": 100}}`

	// fazer o parser do json
	v, err := parse(jsonAccount)
	if err != nil {
		return
	}

	switch v.(type) {
	case *account.Account:
		acc := v.(*account.Account) // convertendo para o tipo account
		if ok := account.Save(acc); ok == false {
			acc.Violations = append(acc.Violations, "account-already-initialized")
			fmt.Println(acc)
			return
		}

	default:
		fmt.Println("Deu ruim...")
	}

}

func parse(s string) (interface{}, error) {
	r := regexp.MustCompile(".+account.+")
	if r.MatchString(s) {
		account, _ := parseAccount(s)
		return account, nil
	}

	return nil, nil
}

func parseAccount(s string) (*account.Account, error) {
	a := account.NewAccount()

	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		return nil, errors.New("Erro ao converter para o tipo Account")
	}

	return &a, nil
}
