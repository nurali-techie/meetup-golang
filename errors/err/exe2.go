package err

import (
	"fmt"
	"log"
)

type LowBalanceError struct {
	err            string
	currentBalance int
}

func (e LowBalanceError) Error() string {
	return fmt.Sprintf("%s, %d", e.err, e.currentBalance)
}

type RoundError struct {
	factor int
}

func (e RoundError) Error() string {
	return fmt.Sprintf("factor %d needed", e.factor)
}

type Account struct {
	Name    string
	Balance int
}

func (a *Account) String() string {
	return fmt.Sprintf("%s has %d", a.Name, a.Balance)
}

func (a *Account) Withdrawal(amount int) error {
	if amount > a.Balance {
		return LowBalanceError{err: "insufficient fund", currentBalance: a.Balance}
	}

	if amount%100 != 0 {
		return RoundError{100}
	}

	a.Balance = a.Balance - amount
	return nil
}

func CustomErrorDemo() {
	fmt.Println("*** Custom Error Demo ***")

	a := &Account{Name: "hemal", Balance: 105}
	fmt.Println("account created, a=", a)

	err := a.Withdrawal(200)
	if lerr, ok := err.(LowBalanceError); ok {
		err = a.Withdrawal(lerr.currentBalance)
		if rerr, ok := err.(RoundError); ok {
			amt := lerr.currentBalance - (lerr.currentBalance % rerr.factor)
			err = a.Withdrawal(amt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("final, a=", a)
}
