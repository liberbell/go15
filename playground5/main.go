package main

import "fmt"

func main() {
	// var s1 []int
	// s1 := make([]int, 4)
	// s1 := []int{0, 1, 2, 3}
	// s1[1] = 10
	// s1Append := []int{4, 5, 6}
	// s1 = append(s1, s1Append...)
	// fmt.Println("value of s1[1]: ", s1[1])

	// s1 := []int{0, 1, 2, 3}
	// s1Copy := make([]int, len(s1))
	// e1 := copy(s1Copy, s1)
	// fmt.Println("length of s1Copy: ", len(s1Copy))
	// fmt.Println("values: ", s1Copy)
	// fmt.Println("number of elements copied: ", e1)

	sl := []int{0, 1, 2, 3, 4, 5}
	sl1 := sl[2:4]
	sl2 := sl[:5]
	sl3 := sl[3:]
	fmt.Println("values in sl1: ", sl1)
	fmt.Println("valeus in sl2: ", sl2)
	fmt.Println("values in sl3: ", sl3)
}
