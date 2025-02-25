package main

import (
	"fmt"
)

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator mst not be zero")
// 	}
// 	return a / b, nil
// }

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator mst not be zero"
	}
	return a / b, ""
}

func main() {
	fmt.Println("started Error handling in the functions")
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Multiplication of two numbers is :", ans)

	ans, _ := divide(10, 0)
	fmt.Println("Multiplication of two numbers is :", ans)

}
