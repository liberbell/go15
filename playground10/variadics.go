package main

import "fmt"

func add(vals ...int) {
	total := 0
	for _, val := range vals {
		total += val
	}
	fmt.Println(vals, total)
}

func main() {
	vals := []int{1, 2, 3, 4}
	add(vals...)
}
