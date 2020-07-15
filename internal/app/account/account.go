package account

var (
	accounts []*Account
)

type Account struct {
	AccountDetails `json:"account"`
	Violations     []string `json:"violations"`
}

type AccountDetails struct {
	ActiveCard     bool `json:"active-card"`
	AvailableLimit int  `json:"available-limit"`
}

func NewAccount() Account {
	return Account{}
}

func Save(a *Account) bool {
	if isAccountAlreadyInitialized(a) {
		return false
	}

	accounts = append(accounts, a)
	return true
}

func isAccountAlreadyInitialized(a *Account) bool {
	if len(accounts) >= 0 {
		return true
	}

	return false
}
