package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(findOs())
	fmt.Println(DayOfWeek())
	fmt.Println(SwitchNoCondition())
}

func findOs() string {
	var result string
	switch os := runtime.GOOS; os {
	case "darwin":
		result = "OS X"
	case "linux":
		result = "Linux"
	default:
		result = os
	}
	return result
}

func DayOfWeek() string {
	var result string
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		result = "Today"
	default:
		result = "Too far away"
	}
	return result
}

func SwitchNoCondition() string {
	var result string
	t := time.Now()
	switch {
	case t.Hour() < 12:
		result = "Good Morning"
	case t.Hour() < 17:
		result = "Good Afternoon"
	default:
		result = "Good Evening"
	}
	return result
}
