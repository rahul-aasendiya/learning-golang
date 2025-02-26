package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId" `
	Id        int    `json:"id" `
	Title     string `json:"title" `
	Completed bool   `json:"completed" `
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error geting", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response")
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding: ", err)
		return
	}
	fmt.Println("Todo: ", todo)
}
func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "Rahul",
		Completed: true,
	}
	// convert the Todo structure to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}
	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	myURL := "https://jsonplaceholder.typicode.com/todos"

	// send post request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()

	// data, _ := ioutil.ReadAll((res.Body))
	// fmt.Println("Response : ", string(data))

	fmt.Println("Response status : ", res.Status)
}
func performUpdateRequest() {
	todo := Todo{
		UserId:    2233,
		Title:     "Rahul kumar",
		Completed: false,
	}

	// convert the Todo structure to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	// Create PUT Request
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request : ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Resonse : ", string(data))
	fmt.Println("Response status : ", res.Status)
}

func performDeleteRequest() {
	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	// Create DELETE Request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error creating PUT request : ", err)
		return
	}

	// send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response status : ", res.Status)

}
func main() {
	fmt.Println("We are learning CRUD in Golang")
	performGetRequest()
	performPostRequest()
	performUpdateRequest()
	performDeleteRequest()

}
