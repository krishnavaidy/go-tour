package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	// should be avoided - harms readability
	return
}

func main() {
	fmt.Println(split(17))
}
