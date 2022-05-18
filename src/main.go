package main

import (
	"fmt"

	ao "local/alura-BasicGo/pkg/accountOwner"
	a "local/alura-BasicGo/pkg/accounts"
)

func PayFees(account IAccount, amount float64) {
	account.Withdraw(amount)
}

type IAccount interface {
	Withdraw(amount float64)
}

func main() {
	accountOwner := ao.AccountOwner{
		Name: "John Doe",
		CPF:  "12313123",
		Job:  "teacher",
	}

	accountA := a.Account{
		Owner:   accountOwner,
		Agency:  589,
		Account: 123456,
		Balance: 112.0,
	}

	accountB := a.Account{ao.AccountOwner{"Some body", "123123123", "singer"}, 589, 123, 100}

	var accountC *a.Account
	accountC = new(a.Account)
	accountC.Balance = 123.4

	accounts := []a.Account{accountA, accountB, *accountC}
	for _, account := range accounts {
		fmt.Println(account)
	}

	widthdrawAmount := 10.
	accountA.Withdraw(widthdrawAmount)
	accountB.Withdraw(0.0)

	depositAmount := 500.0
	accountA.Deposit(depositAmount)

	accountA.Transfer(100, &accountB)

	PayFees(&accountA, 100)

	accountD := a.HoldingAccount{}
	accountD.Deposit(depositAmount)
	accountD.Withdraw(widthdrawAmount)
	PayFees(&accountD, 100)
}
