package main

import "fmt"

type Food struct {
	Nutrition() string
	FoodType() string
}

type apple strust {
}

func (a Apple) Nutrition() string {
	return "Apples are carb heavy!"
}

func (a Apple) FoodType() string {
	return "Apples are fruit."
}

type Celery struct {
}

func (c Celery) Nutrition() string {
	return "Celery has zero everything!"
}

func (c Celery) FoodType() string {
	return "Celery is a vegetable."	
}

func main() {
	foods := []Food{Apple{}, Celery{}}
	for _, f := range foods {
		fmt.Println(f.Nutrition(), f.FoodType())
	}
}