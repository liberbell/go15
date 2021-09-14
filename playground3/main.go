package main

import "fmt"

func main() {
	x := 5
	y := 2

	fmt.Println(x+y, x-y, x*y, x/y, x%y)
	fmt.Println(x == y, x != y, x > y, x < y, x >= 5)
	fmt.Println(x > y && x >= 5)
	fmt.Println(x > y || x < y)
	fmt.Println(!true, !false)

	fmt.Println(2 << 1)
}
