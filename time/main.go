package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("current time : ", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006 , 15:04:05")
	fmt.Println("Formatted time :", formatted)

	layoutStr := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Formatted time:", formatted_time)

	//add a day in cuurent time
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new date:", new_date)
	formatted_new_date := new_date.Format("2006/01/02")
	fmt.Println("formatted_new_date:", formatted_new_date)
}
