package main

import "fmt"

func main() {
	// workday := 3
	// switch workday {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Thuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// default:
	// 	fmt.Println("Take the weekend off")
	// }

	temp := -10
	switch {
	case temp < 0:
		fmt.Println("Below freezing.")
	case temp == 0:
		fmt.Println("At freezing point.")
	default:
		fmt.Println("Above freezing.")
	}
}
