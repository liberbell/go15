package main

import "fmt"

func main() {
	// sl := []int{10, 20, 30, 40}
	// for i, value := range sl {
	// 	fmt.Println(i, value)
	// }
	prodPrice := map[string]int{
		"widget":             75,
		"turbo widget":       100,
		"convertible widget": 150,
	}
	for _, value := range prodPrice {
		fmt.Println(value)
	}
}
