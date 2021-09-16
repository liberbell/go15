package main

import "fmt"

func main() {
	// var s1 []int
	// s1 := make([]int, 4)
	s1 := []int{0, 1, 2, 3}
	fmt.Println("values: ", s1)
	fmt.Println("is zero slice nil?", s1 == nil)
}
