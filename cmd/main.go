package main

import (
	"fmt"

	"github.com/gpraul/authorizer-challange/cmd/adapter"
	"github.com/gpraul/authorizer-challange/cmd/core"
	"github.com/gpraul/authorizer-challange/cmd/parser"
)

func main() {
	jsonAccount := `{"account": {"active-card": true, "available-limit": 100}}`

	// parser do json
	operation, err := parser.Parse(jsonAccount)

	if err != nil {
		fmt.Println(err)
		return
	}

	// salvando a conta em um slice
	account := operation.(core.Account)
	adapter.SaveAccount(account)

	account, err = adapter.FindOneAccount()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Nova conta:", account)

}
