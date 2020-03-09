package main

import "fmt"

// variadic functions can take any amount of inputs like **kwargs
func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for x, y := range nums {
		fmt.Println(x, y)
		total += y
	}

	return total
}

func main() {

	fmt.Println("Sum of number is: ", sum(1, 2, 3, 4, 5, 6))
}
