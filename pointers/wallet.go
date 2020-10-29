package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds error when not enough founds in wallet
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Deposit into wallet
func (w *Wallet) Deposit(val Bitcoin) {
	w.balance += val
}

// Balance of actual wallet
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw of actual wallet
func (w *Wallet) Withdraw(val Bitcoin) error {
	if val > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= val
	return nil
}

// String change bitcoin definition
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
