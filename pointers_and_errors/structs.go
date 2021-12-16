package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrorBalanceIsNegativeWhileWithdraw = errors.New("you can't withdraw amount of money that's greater than your balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorBalanceIsNegativeWhileWithdraw
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
