package main

import "fmt"

type Box struct {
	D float64
	W float64
	H float64
}

func main() {

	b := Box{5, 4, 4}
	fmt.Println(b)
}
