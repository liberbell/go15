package main

import "fmt"

var x int = 20
var y int = 10

func add() int {
	return x + y
}

func main() {
	// x := 20
	// y := 10
	fmt.Println(add())
}
