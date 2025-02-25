package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana, princce"
	parts := strings.Split(data, ",")
	fmt.Println("parts: ", parts)

	str := "one two three four wo two two five"
	count := strings.Count(str, "two")
	fmt.Println("Count: ", count)

	str2 := "  Hello, GO!   "
	trimed := strings.TrimSpace(str2)
	fmt.Println("Trimed: ", trimed)

	str3 := "Rahul"
	str4 := "Aasendiya"

	combined_string := strings.Join([]string{str3, "kumar", str4}, " ")
	fmt.Println("Combined String: ", combined_string)

}
