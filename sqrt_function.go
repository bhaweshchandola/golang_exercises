package main

import "fmt"

func sqrt(x float64) float64 {

	z := float64(1)

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("square root is: ", sqrt(16))
}
