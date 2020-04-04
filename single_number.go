package main

import (
	"fmt"
)

func search_in_array(arr []int, num int) []int {
	for x, y := range arr {
		if y == num {
			arr[len(arr)-1], arr[x] = arr[x], arr[len(arr)-1]
			return arr[:len(arr)-1]
		}
	}

	return arr
}

func main() {
	seen := make([]int, 0)
	inp := [5]int{2, 2, 1, 1, 4}
	fmt.Println(seen, inp)
	for _, y := range inp {
		if len(seen) == 0 {
			seen = append(seen, y)
		} else {
			seen = search_in_array(seen, y)
		}
	}
	fmt.Println("final_answer", seen[0])
}
