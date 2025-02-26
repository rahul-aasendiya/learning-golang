package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("We are learning url in Golang")
	myURL := "https://example.com/path/to/resource?key1=value1&key2=value2"

	fmt.Printf("Type of url: %T\n", myURL)
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}
	fmt.Printf("Type of parsedURL: %T\n", parsedURL)

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)

	// Modify URL Component
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=rahul"

	// Constructing a URLstrin from url object
	newURL := parsedURL.String()

	fmt.Println("New URL: ", newURL)

}
