package main

import "fmt"

func e35() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	(*p).y = 3
	fmt.Println(v)
	fmt.Println(*p)
}
