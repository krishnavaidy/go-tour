package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
	n1 := 0
	n2 := 1
	return func() int {
		sum := n1 + n2
		n1 = n2
		n2 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
