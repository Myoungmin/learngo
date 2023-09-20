package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var NonMoney = errors.New("Can't withdraw you are poor")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return NonMoney
	}
	a.balance -= amount
	return nil
}
