package main

import "fmt"

func avg(s []float64) float64 {
	total := 0.0
	for _, val := range s {
		total += val
	}
	return total / float64(len(s))
}

func main() {
	s := []float64{1.8, 2.2, 2.0}
	fmt.Println(avg(s))
}
