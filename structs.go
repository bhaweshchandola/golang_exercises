package main

import "fmt"

type Ver struct {
	x int
	y int
}

var (
	v1 = Ver{1, 2}
	v2 = Ver{x: 20} // here y is declared as 0
	v3 = Ver{}      //both x and y are zero
)

func main() {

	v := Ver{1, 2}
	fmt.Println("x and y value of the struct", v.x, v.y)
	fmt.Println("v1, v2, v3", v1, v2, v3)
}
