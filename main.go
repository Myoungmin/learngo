package main

import (
	"fmt"

	"github.com/Myoungmin/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "nicolas", Balance: 1000}
	fmt.Println(account)
}
