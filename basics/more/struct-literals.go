package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v2 := Vertex{X: 1}
	fmt.Println(v)
	fmt.Println(v2)
}
