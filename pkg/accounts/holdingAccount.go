package accounts

import (
	"fmt"

	ao "local/alura-BasicGo/pkg/accountOwner"
)

type HoldingAccount struct {
	Owner                      ao.AccountOwner
	Agency, Account, Operation int
	Balance                    float64
}

func (a *HoldingAccount) Withdraw(amount float64) {
	allowWithdraw := amount <= a.Balance
	if amount <= 0.0 {
		fmt.Println("Error: Amount must be greater than 0")
		return
	}

	if allowWithdraw {
		oldBalance := a.Balance

		a.Balance -= amount

		printTransactionInfo(amount, oldBalance, a.Balance)
		return
	}

	fmt.Println("Error: Balance insufficient")
}

func (a *HoldingAccount) Deposit(amount float64) {
	if amount <= 0.0 {
		fmt.Println("Error: Amount must be greater than 0")
		return
	}

	oldBalance := a.Balance

	a.Balance += amount

	printTransactionInfo(amount, oldBalance, a.Balance)
}
