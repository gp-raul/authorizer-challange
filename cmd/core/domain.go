package core

import "time"

// Account struct
type Account struct {
	AccountDetail `json:"account"`
	Violation     []string `json:"violations"`
}

// AccountDetail struct
type AccountDetail struct {
	ActiveCard     bool `json:"active-card"`
	AvailableLimit int  `json:"available-limit"`
}

// Transaction struct
type Transaction struct {
	TransactionDetail `json:"transaction"`
}

// TransactionDetail struct
type TransactionDetail struct {
	Merchant string    `json:"merchant"`
	Amount   int       `json:"amount"`
	Time     time.Time `json:"time"`
}
