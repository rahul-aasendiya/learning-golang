package main

import (
	"fmt"
)

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func multiplication(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learning function in GOLANG...")
	simpleFunction()
	add := add(3, 4)
	mult := multiplication(3, 4)
	fmt.Println("Add of two numbers is :", add)
	fmt.Println("Multiplication of two numbers is :", mult)
}
