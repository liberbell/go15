package main

import "fmt"

func alter(s string) string {
	s[5] = " "
	return string(s)
}

func main() {
	msg := "Hello World"
	fmt.Println(alter(msg))
}
