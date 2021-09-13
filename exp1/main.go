package main

import "fmt"

func alter(s string) string {
	s[5] = " "
}

func main() {
	msg := "Hello World"
	fmt.Println(alter(msg))
}
