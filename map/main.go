package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are learning map in Golang")
	studentGrades := make(map[string]int)

	studentGrades["rahul"] = 30
	studentGrades["amit"] = 60
	studentGrades["suresh"] = 70
	studentGrades["anil"] = 80

	fmt.Println("Marks of anil:", studentGrades["anil"])
	studentGrades["anil"] = 20
	fmt.Println("Marks of anil:", studentGrades["anil"])

	delete(studentGrades, "anil")
	fmt.Println("Marks of anil:", studentGrades["anil"])

	// Checking f a key exists
	grades, exists := studentGrades["arjun"]
	fmt.Println("Grades of arjun : ", grades)
	fmt.Println("Arjun exists : ", exists)

	// Checking f a key exists
	Grades, Exists := studentGrades["amit"]
	fmt.Println("Grades of amit : ", Grades)
	fmt.Println("amit exists : ", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"Amit":  10,
		"Arjun": 40,
		"Anita": 67,
	}

	for index, value := range person {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}
}
