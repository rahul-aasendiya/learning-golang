package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are learning Array in Golang")
	// var name [5]string
	// name[0] = "Prince"
	// name[2] = "Akash"
	// fmt.Println("name of person is :", name)

	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Number is :", numbers)

	// fmt.Println("Length of numbers array is :", len(numbers))

	// fmt.Println("Value of name at 2 index is :", name[2])

	// for numeric types (int, float, etc.) the value is 0. for string, it is an empty string ("").
	// For boolean types, it is false , and for pointers or complex types it is nil.
	var price [5]int
	fmt.Println("Price is :", price)
	fmt.Printf("price array is %q\n", price)

}
