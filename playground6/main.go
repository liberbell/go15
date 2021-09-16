package main

import (
	"fmt"
)

func main() {
	var prodPrice map[string]int
	fmt.Println(prodPrice)
	// prodPrice["widget"] = 100

	tempPrice := make(map[string]int)
	tempPrice["convertible widget"] = 150
	prodPrice = tempPrice
	prodPrice["widget"] = 100
	fmt.Println(prodPrice)
}
