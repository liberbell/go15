package main

import "fmt"

type Box struct {
	D float64
	W float64
	H float64
}

func main() {

	b := Box{D: 5, W: 4, H: 3}
	// b.D = 6
	ptr := &b
	(*ptr).D = 7
	fmt.Println(b)
}
