package main

import (
	"fmt"
	"time"
)

func main() {
	// Standard datetime formate for go specially
	// For 24 hours formate: 02-01-2006, 15:04:05
	// For 24 hours formate: 02-01-2006, Monday
	// For 12 hours formate: 02-01-2006, 3:04 PM
	currentTime := time.Now()
	fmt.Println("Curent Time: ", currentTime)
	fmt.Printf("Type of curentTime %T\n ", currentTime)

	formatted := currentTime.Format("02-01-2006, 3:04 PM")
	fmt.Println("Curent Date: ", formatted)

	// layout_str := "2006-01-02"
	// dateStr := "2013-11-25"
	layout_str := "02/01/2006"
	dateStr := "25/11/2013"
	formated_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formated Time : ", formated_time)

	// Add 1 more day in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	formated_new_date := new_date.Format("02-01-2006 Monday")
	fmt.Println("Formated New Date : ", formated_new_date)

}
