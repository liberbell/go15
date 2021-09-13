package main

import (
	"fmt"
)

func main() {
	// const pi = 3.14
	// pi = 3.14159
	// fmt.Println(math.Pi)
	// fmt.Println(pi)

	const myConst int = 1
	fmt.Printf("type: %T, value: %v\n", float64(myConst)+1.1, float64(myConst)+1.1)
}
