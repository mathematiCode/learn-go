package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	// this defines how to handle using %s when printing Bitcoin types. It doesn't need to be called explicitly because it will be used whenever a %s is used on something of type Bitcoin.
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance-amount >= 0 {
		w.balance -= amount
		return nil
	} else {
		return errors.New("insufficient funds")
	}
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("wallet pointer is %p and balance pointer is %p \n", &w, &w.balance)
	return w.balance
}
