package main

import "fmt"

type Box struct {
	D float64
	W float64
	H float64
}

func (b *Box) volume() float64 {
	return b.D * b.W * b.H
}

func main() {
	b := Box{D: 5, W: 4, H: 3}
	fmt.Println(b.volume())
}
