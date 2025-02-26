package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, world!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("say hello function ended successfully.")
}
func sayHey() {
	fmt.Println("Hey rahul!")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("We are learning go routine in Golang")
	go sayHello()
	go sayHey()

	// Wail for a moment to allow the go routine to finish
	time.Sleep(900 * time.Millisecond)
}
