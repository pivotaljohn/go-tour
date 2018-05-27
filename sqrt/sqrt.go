package main

import (
	"fmt"
)

func abs(x float64) (a float64) {
	if x >= 0 {
		a = x
	} else {
		a = -x
	}
	return
}

func sqrt(x float64) (float64, int) {
	z := float64(1)
	delta := 1.0
	i := 1
	for newZ := float64(0); delta > 0.00001; {
		newZ = z - (z*z-x)/(2*z)
		delta = abs(z - newZ)
		fmt.Printf("Iteration %v: z = %v; newZ = %v; delta = %v\n", i, z, newZ, delta)
		i++
		z = newZ
	}
	return z, i
}

func main() {
	x := 49.0
	sr, i := sqrt(x)
	fmt.Printf("The square root of %v = %v (after %v iterations)", x, sr, i)
}
