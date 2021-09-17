package main

import "fmt"

func main() {

	val := 123
	var ptr *int = &val
	fmt.Println(ptr)
}
