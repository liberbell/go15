package main

import "fmt"

func main() {
	if score := 99; score < 100 {
		fmt.Println("Oops")
	} else if score > 100 && score < 1000 {
		fmt.Println("good")
	} else {
		fmt.Println("Wow")
	}
}
