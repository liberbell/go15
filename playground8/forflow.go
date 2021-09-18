package main

import "fmt"

func main() {
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 0; i < 5; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	i := 0

outterlabel:
	for i < 5 {
		if i == 2 {
			i++
			goto outterlabel
		}
		fmt.Println(i)
		i++
	}

}
