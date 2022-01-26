package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner:owner, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't Withdraw")
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, " has ", a.balance)
}