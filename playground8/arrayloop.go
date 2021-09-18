package main

import "fmt"

func main() {
	sl := []int{10, 20, 30, 40}
	for i, value := range sl {
		fmt.Println(i, value)
	}
}
