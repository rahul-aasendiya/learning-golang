package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are learning Array in Golang")
	var name [5]string
	name[0] = "Prince"
	name[2] = "Akash"
	fmt.Println("name of person is :", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is :", numbers)

	fmt.Println("Length of numbers array is :", len(numbers))

	fmt.Println("Value of name at 2 index is :", name[2])

}
