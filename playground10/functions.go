package main

import "fmt"

// var x int = 20
// var y int = 10

func add(x, y, z int) int {
	return x + y + z
}

func main() {
	// x := 20
	// y := 10
	// z := 5
	s := []int{10, 20, 5}
	fmt.Println(add(s))
}
