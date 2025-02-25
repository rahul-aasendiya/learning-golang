package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are learning For Loop in Golang")
	// for i := 0; i < 4; i++ {
	// 	fmt.Println("Number is :", i)
	// }

	// counter := 0
	// for {
	// 	fmt.Println("Infinite Loop")
	// 	counter++
	// 	if counter == 3 {
	// 		break
	// 	}
	// }
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	for index, value := range numbers {
		fmt.Printf("Index of numbers is %d, and value is %d\n", index, value)
	}

	data := "Hello, world!"
	for index, char := range data {
		fmt.Printf("Index of data is %d, and value is %c\n", index, char)
	}
}
