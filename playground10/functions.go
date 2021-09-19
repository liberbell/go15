package main

import "fmt"

func add() int {
	return x + y
}

func main() {
	x := 20
	y := 10
	fmt.Println(add())
}
