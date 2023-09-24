// package accounts

// import (
// 	"fmt"

// 	"github.com/Myoungmin/learngo/accounts"
// )

// func main_accounts() {
// 	account := accounts.NewAccount("nico")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.Balance(), account.Owner())
// 	fmt.Println(account)
// }

package main

import (
	"fmt"

	"github.com/Myoungmin/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	definition, err = dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
