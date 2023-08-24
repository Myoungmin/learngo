package main

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (length int, uppercase string) {
	// function이 값을 return하고 나면 실행
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	//fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper("Myoungmin")
	// fmt.Println(totalLength, upperName)

	// totalLength, _ := lenAndUpper("Myoungmin")
	// fmt.Println(totalLength)

	//repeatMe("nico", "lynn", "dal", "marl", "flynn")

	totalLength, up := lenAndUpper("Myoungmin")
	fmt.Println(totalLength, up)
}
