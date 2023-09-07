package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch {
	case age < 19:
		return false
	case age >= 19:
		return true
	}
	return false
}
func main() {
	fmt.Println(canIDrink(16))
}
