package main

import (
	"fmt"
)

func main() {
	// const pi = 3.14
	// pi = 3.14159
	// fmt.Println(math.Pi)
	// fmt.Println(pi)

	const myConst = setMe()
	fmt.Printf("type: %T, value: %v\n", myConst+1.1, myConst+1.1)
}

func setMe() int {
	return 1
}
