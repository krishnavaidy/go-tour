package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var table [][]uint8
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			table[i][j] = (i + j) / 2
		}
	}
}

func main() {
	pic.Show(Pic)
}
