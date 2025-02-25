package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("We are learning data conversion in golang.")
	var num int = 42

	fmt.Println("Number is: ", num)
	fmt.Printf("Type of num is %T\n", num)

	// Convert integer to float
	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("Number is: ", data)
	fmt.Printf("Type of num is %T\n", data)

	// Convert integer to string
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("num is: ", str)
	fmt.Printf("Type of num is %T\n", str)

	// Convert  string to integer
	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("number_int is: ", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	// Convert  string to float
	num_string := "12.34"
	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("number_float is: ", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)

}
