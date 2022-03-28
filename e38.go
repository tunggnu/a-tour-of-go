package main

import "fmt"

func e38() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	s = primes[2:]
	fmt.Println(s)
}
