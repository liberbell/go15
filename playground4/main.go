package main

import "fmt"

func main() {
	i := 10
	for i := 0; i < 5; i++ {
		fmt.Println("The value of i inside the for loop: ", i)
	}
	fmt.Println("The value of i outside the for loop: ", i)

	for i < 15 {
		fmt.Println("using the value of i initialized outside for loop: ", i)
		i++
	}
	fmt.Println("value of i outside the for loop did change: ", i)

	j := 1
	if j < 5 {
		fmt.Println("value of j inside the if statement: ", j)
	}
}
