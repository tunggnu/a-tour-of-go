package main

import "fmt"

func e20() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
