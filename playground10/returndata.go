package main

import "fmt"

func add(x, y int) (a, b int, c bool) {
	return x + y
}

func main() {
	x := 20
	y := 10
	fmt.Println(add(x, y))

}
