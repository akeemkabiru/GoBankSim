package bank_operations

type Account struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	InitialDeposit string `json:"initial_deposit"`
}

func NewAccount(firstName, lastName, initialDeposit string) Account {
	return Account{FirstName: firstName, LastName: lastName, InitialDeposit: initialDeposit}
}
