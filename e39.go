package main

import "fmt"

func e39() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XYZ"
	fmt.Println(a, b)
	fmt.Println(names)
}
