package main

import (
	"github.com/gpraul/authorizer-challange/cmd/parser"
)

func main() {
	jsonAccount := `{"account": {"active-card": true, "available-limit": 100}}`

	parser.Parse(jsonAccount)
}
