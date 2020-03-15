package main

import "fmt"

func main() {

	i, j := 20, 200

	p := &i
	fmt.Println("pointer and address", *p, p)

	p = &j
}
