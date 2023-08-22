package main

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	//fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("Myoungmin")
	fmt.Println(totalLength, upperName)
}
