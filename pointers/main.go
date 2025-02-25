package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num * 20

}
func main() {
	fmt.Println("We are learning pointer in golang.")
	num := "Rahul"
	ptr := &num

	fmt.Println("Num has value: ", num)
	fmt.Println("Pointer contain: ", ptr)
	fmt.Println("Data contains through pointer: ", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned.")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value contains : ", value)
}
