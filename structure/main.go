package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	House int
	Area  string
	State string
}

type Contact struct {
	Email string
	Phone string
	Fax   string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	fmt.Println("We are learning struture in Golang")
	var prince Person
	// fmt.Println("Prince person : ", prince)
	prince.FirstName = "Prince"
	prince.LastName = "Agarwal"
	prince.Age = 24
	fmt.Println("Prince person : ", prince)

	person1 := Person{
		FirstName: "Akash",
		LastName:  "Singh",
		Age:       25,
	}
	fmt.Println("Person 1 : ", person1)

	// person2 := Person{
	// 	FirstName: "Susmita",
	// 	LastName:  "sen",
	// 	Age:       40,
	// }
	// fmt.Println("Person 2 : ", person2)
	// fmt.Println("age of prince is : ", prince.Age)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Prince",
		LastName:  "Agarwal",
		Age:       26,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "8585859595"

	employee1.Person_Address = Address{
		House: 12,
		Area:  "Ranchi",
		State: "Jharkhand",
	}
	employee1.Person_Contact.Fax = "Fax@1984616"

	fmt.Println("Employee 1 : ", employee1)

}
