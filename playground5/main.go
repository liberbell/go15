package main

import "fmt"

func main() {
	// var s1 []int
	s1 := make([]int, 4)
	fmt.Println("values: ", s1)
	fmt.Println("is zero slice nil?", s1 == nil)
}
