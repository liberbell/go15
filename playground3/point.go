package main

import "fmt"

func main() {
	x := "hello"
	y := &x

	fmt.Println(x, y)
	fmt.Println(*y)
	*y = "goodbye"
	fmt.Println(x, y)
}
