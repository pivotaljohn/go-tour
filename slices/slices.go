package main

import (
	"golang.org/x/tour/pic"
	// "math"
)

func graph(x, y int) uint8 {
	//	return uint8(255*math.Sin(float64(x))+255*math.Cos(float64(y)))
	return uint8(255 - (y-x)/2)
}

func Pic(dx, dy int) [][]uint8 {
	img := [][]uint8{}

	for i := 0; i < dy; i++ {
		row := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			row[j] = graph(i, j)
		}
		img = append(img, row)
	}

	return img
}

func main() {
	pic.Show(Pic)
}
