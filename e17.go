package main

import "fmt"

const Pi = 3.14

func e17() {
	const World = "Thế giới"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
