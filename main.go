package main

import (
	"fmt"

	"github.com/Myoungmin/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
