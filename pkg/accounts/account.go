package accounts

import (
	"fmt"
	"strconv"

	ao "local/alura-BasicGo/pkg/accountOwner"
)

type Account struct {
	Owner   ao.AccountOwner
	Agency  int
	Account int
	Balance float64
}

func (a *Account) Withdraw(amount float64) {
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

func (a *Account) Deposit(amount float64) {
	if amount <= 0.0 {
		fmt.Println("Error: Amount must be greater than 0")
		return
	}

	oldBalance := a.Balance

	a.Balance += amount

	printTransactionInfo(amount, oldBalance, a.Balance)
}

func (a *Account) Transfer(amount float64, accountToTransfer *Account) {
	if amount <= 0.0 {
		fmt.Println("Error: Amount must be greater than 0")
		return
	}

	if amount <= a.Balance {
		oldBalance := a.Balance

		a.Balance -= amount
		accountToTransfer.Balance += amount

		printTransactionInfo(amount, oldBalance, a.Balance)
		return
	}
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func printTransactionInfo(amount float64, oldBalance float64, newBalance float64) {
	oldBalanceString := strconv.FormatFloat(oldBalance, 'f', 2, 64)
	amountString := strconv.FormatFloat(amount, 'f', 2, 64)
	newBalanceString := strconv.FormatFloat(newBalance, 'f', 2, 64)

	fmt.Println("****************************************************************")
	fmt.Println("Transaction happened sucessfully.")
	fmt.Println("Old Balance:" + oldBalanceString)
	fmt.Println("Amount:" + amountString)
	fmt.Println("New Balance:" + newBalanceString)
	fmt.Println("****************************************************************")
}
