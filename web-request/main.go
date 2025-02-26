package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("We are learning web request in Golang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting from Get response ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of respones: %T\n", res)

	// Read the response body

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response ", err)
		return
	}
	fmt.Println("Byte Response: ", data)
	fmt.Println("Response: ", string(data))

}
