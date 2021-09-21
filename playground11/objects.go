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

type Sphere struct {
	radius float64
}

func (sp *Sphere) volume() float64 {
	return ((4.0 / 3.0) * 3.14 * (sp.radius * sp.radius * sp.radius))
}

func main() {
	a
}
