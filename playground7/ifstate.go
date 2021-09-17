package main

import "fmt"

func main() {
	temp := 10
	if temp < 0 {
		fmt.Println("Below freezing!")
	} else if temp == 0 {
		fmt.Println("At freezing point!")
	} else {
		fmt.Println("Above freezing!")
	}
}
