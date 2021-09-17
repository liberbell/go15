package main

import "fmt"

func main() {

	// val := 123
	// var ptr *int = &val
	// fmt.Println(ptr, *ptr)
	var ptr *int = new(int)
	fmt.Println(ptr, *ptr)
	*ptr = 234
	fmt.Println(ptr, *ptr)

	val := 123
	fmt.Println(ptr, *ptr)
}
