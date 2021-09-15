package main

import "fmt"

func main() {
	// var arr [4]int
	arr := [4]int{1, 2, 3, 4}
	arr[3] = 10
	fmt.Println("value of array[1]: ", arr[1])
	fmt.Println("length of array: ", len(arr))
	fmt.Println("values: ", arr)
}
