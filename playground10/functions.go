package main

import "fmt"

// var x int = 20
// var y int = 10

func add(x []int) int {
	total := 0
	for _, val := range x {
		total += val
	}
	return total
}

func multiply(x, y *int) int {
	return *x * *y
}

func main() {
	// x := 20
	// y := 10
	// z := 5
	// s := []int{10, 20, 5}
	x := 5
	y := 2
	fmt.Println(multiply(&x, &y))
}
