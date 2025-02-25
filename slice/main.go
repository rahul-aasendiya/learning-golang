package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are learning slice in Golang")
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 2, 3, 5, 546, 456, 345, 6, 34, 43, 5, 2345, 234, 5, 234, 5234, 5, 2435, 23)
	// fmt.Println("Numbers : ", numbers)
	// fmt.Printf("Number has data type: %T\n", numbers)
	// fmt.Println("Length : ", len(numbers))

	numbers := make([]int, 3, 5)
	numbers = append(numbers, 2)
	numbers = append(numbers, 90)
	numbers = append(numbers, 100)
	fmt.Println("Numbers Slice:", numbers)
	fmt.Println("Numbers Length:", len(numbers))
	fmt.Println("Numbers Capacity:", cap(numbers))

	stock := make([]int, 0)
	fmt.Println("Stock Slice:", stock)
	fmt.Println("Stock Length:", len(stock))
	fmt.Println("Stock Capacity:", cap(stock))

}
