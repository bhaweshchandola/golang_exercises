package main

import "fmt"

// naked return as no return values are specified but the named return in the top are returned
func named_ret(sum int) (x, y int) {
	x = sum + 1
	y = sum + 10
	return
}

func main() {
	x, y := named_ret(20)
	fmt.Println("x,y", x, y)
}
