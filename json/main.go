package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name" `
	Age     int    `json:"age" `
	IsAdult bool   `json:"is_adult" `
}

func main() {
	fmt.Println("We are learning json in Golang")
	person := Person{Name: "Jhon", Age: 43, IsAdult: true}
	// fmt.Println("Person Data is: ", person)

	// convert person into JSON Encodign (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling ", err)
		return
	}
	fmt.Println("Person Data is: ", string(jsonData))

	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling ", err)
		return
	}
	fmt.Println("Person data is after unmarshalling: ", personData)
}
