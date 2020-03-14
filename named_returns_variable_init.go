package main

import "fmt"

// naked return as no return values are specified but the named return in the top are returned
func named_ret(sum int) (x, y int) {
	x = sum + 1
	y = sum + 10
	return
}

// declaration of variables can be doe at package/function level or inside main
var c, python, java bool

// variable with initializers, in this if 'var' keyword is used type can be skipped
var c1, python1 = 1, 2

func main() {
	x, y := named_ret(20)
	fmt.Println("x,y", x, y)
	var i bool
	fmt.Println(c, python, java, i)
	var j, k, l = true, false, "not here"
	fmt.Println(c1, python1, j, k, l)
}
