package main

import (
	"fmt"
)

func search_in_array(arr []int, num int) bool {
	for _, y := range arr {
		if y == num {
			return true
		}
	}

	return false
}

func main() {
	seen := make([]int, 0)
	inp := [5]int{2, 2, 1, 1, 4}
	fmt.Println(seen, inp)
	seen = append(seen, 4)
	for _, y := range inp {
		fmt.Println(search_in_array(seen, y))
	}
}
