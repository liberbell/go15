package main

import "fmt"

func main() {
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
