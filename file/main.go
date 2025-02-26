package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("We are learning file in Golang")
	// Creating file and add content into file
	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file: ", err)
			return
		}
		defer file.Close()
		content := "Hello world! rahul aasendiya"
		byte, errors := io.WriteString(file, content+"\n")
		fmt.Println("Byte written in file is: ", byte)
		if errors != nil {
			fmt.Println("Error while writing file: ", errors)
			return
		}
		fmt.Println("successfully created fle")
	*/

	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while creating file: ", err)
			return
		}
		// Create a buffer to read the file content
		buffer := make([]byte, 1024)
		// Read the file content into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file", err)
				return
			}

			// Process the read content
			fmt.Println(string(buffer[:n]))

		}
		defer file.Close()
	*/

	// Read te entire file into a byte slice
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	fmt.Println(string(content))

}
