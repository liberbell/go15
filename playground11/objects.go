package main

type Shape interface {
	volumes() float64
}

type Cube struct {
	depth  float64
	width  float64
	height float64
}

func (cu *Cube) volume() float64 {
	return cu.depth * cu.width * cu.height
}

func main() {
	a
}
