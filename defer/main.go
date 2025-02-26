package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("We are learning defer in Golang")
	fmt.Println("start")
	data := add(5, 6)
	defer fmt.Println("Data is : ", data)
	defer fmt.Println("mid")
	fmt.Println("end")
}
