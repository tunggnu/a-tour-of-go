package main

import "fmt"

func e40() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	t := r[1:6]
	t[0] = true
	fmt.Println(r)
	fmt.Println(t)

	s := [](struct {
		i int
		b bool
	}){
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
