package main

import (
	"fmt"
	"time"
)

func e28() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday - today {
	case 0:
		fmt.Println("Today.")
	case 1:
		fmt.Println("Tomorrow.")
	case 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
