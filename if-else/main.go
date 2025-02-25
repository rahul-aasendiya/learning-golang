package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are learning if else in Golang")
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	z := 10
	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than 10")
	} else {
		fmt.Println("z is smaller than 5")
	}

	y := 10

	if x > 5 && y > 5 {
		fmt.Println("Hey how are you?")
	} else {
		fmt.Println("Learning go programming")
	}
}
