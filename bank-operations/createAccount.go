package bank_operations

import "time"

type Account struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	InitialDeposit string    `json:"initial_deposit"`
	CreatedAt      time.Time `json:"created_at"`
	AccountNumber  string    `json:"account_number"`
}

func NewAccount(firstName, lastName, initialDeposit, accountNumber string) Account {
	return Account{FirstName: firstName, LastName: lastName, InitialDeposit: initialDeposit, CreatedAt: time.Now(), AccountNumber: accountNumber}
}
