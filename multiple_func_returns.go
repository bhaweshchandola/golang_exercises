package main

import "fmt"

func multi_return(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(multi_return("hey", "boy"))
}
