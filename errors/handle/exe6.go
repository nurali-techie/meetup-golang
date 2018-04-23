package handle

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
		// return LowBalanceError{err: "insufficient fund", currentBalance: a.Balance}
		panic(LowBalanceError{err: "insufficient fund", currentBalance: a.Balance})
	}

	if amount%100 != 0 {
		// return RoundError{100}
		panic(RoundError{100})
	}

	a.Balance = a.Balance - amount
	return nil
}

func AllInOneDemo() {
	fmt.Println("*** All In One Demo ***")

	a := &Account{Name: "hemal", Balance: 105}
	fmt.Println("account created, a=", a)

	// err := a.Withdrawal(200)	// panic
	err := callWithdral(a, 200) // panic to error
	if lerr, ok := err.(LowBalanceError); ok {
		// err = a.Withdrawal(lerr.currentBalance)	// panic
		err = callWithdral(a, lerr.currentBalance) // panic to error
		if rerr, ok := err.(RoundError); ok {
			amt := lerr.currentBalance - (lerr.currentBalance % rerr.factor)
			err = callWithdral(a, amt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("final, a=", a)
}

func callWithdral(a *Account, amt int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
		}
	}()
	err = a.Withdrawal(amt)
	return err
}
