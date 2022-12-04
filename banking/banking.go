package banking

import "errors"

type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	accoout := Account{owner: owner, balance: 0}
	return &accoout
}

// Deposit x amount on your account
// a Account는 리시버이다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	// error는 nil이나 error 형태가 될 수 있다.
	return nil
}

func (a Account) Balance() int {
	return a.balance
}
