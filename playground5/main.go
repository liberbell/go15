package main

import "fmt"

func main() {
	// var s1 []int
	// s1 := make([]int, 4)
	s1 := []int{0, 1, 2, 3}
	// s1[1] = 10
	s1 = append(s1, 4, 5, 6)
	// fmt.Println("value of s1[1]: ", s1[1])
	fmt.Println("values: ", s1)
	fmt.Println("is zero slice nil?", s1 == nil)
	fmt.Println("length of s1: ", len(s1))
}
