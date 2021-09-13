package main

import (
	"fmt"
)

func main() {
	// const pi = 3.14
	// pi = 3.14159
	// fmt.Println(math.Pi)
	// fmt.Println(pi)

	const myConst = 1
	fmt.Printf("type: %T, value: %v", myConst+1.1, myConst+1.1)
}
