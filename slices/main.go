package main

import "fmt"

func main() {
	var holidays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	var workDays = holidays[0:5]
	holidays = holidays[5:7]
	var weekDays = append(workDays, holidays...)
	fmt.Println(weekDays)
}
