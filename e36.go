package main

import "fmt"

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{2, 1}
)

func e36() {
	fmt.Println(v1, p, v2, v3)
}
