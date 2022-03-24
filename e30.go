package main

import "fmt"

func e30() {
	defer fmt.Println("hello")

	fmt.Println("world")
}
