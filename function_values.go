package main

import "fmt"

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	new_func := func(x, y float64) float64 {
		return x + y
	}

	fmt.Println(new_func(20, 10))

	fmt.Println(compute(new_func))
}
