package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// 조건을 체크하기 전에 variable을 만들 수 있다
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
