package main

import "fmt"

func main() {
	//const name string = "nico"

	//var name string = "nico"
	// 아래는 위와 동일한 코드이다.
	// Go가 type을 찾아준다.
	name := "nico"

	name = "lynn"
	fmt.Println(name)
}
