package main

import (
	"fmt"

	"github.com/Myoungmin/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20)
	fmt.Println(account.Balance())
}
