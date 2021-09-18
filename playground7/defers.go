package main

import "fmt"

func main() {
	defer fmt.Println("second")

	fmt.Println("first")
}
