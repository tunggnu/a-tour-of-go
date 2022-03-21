package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func e08() {
	var a, b = swap("hello", "world")

	fmt.Println(a, b)
}
