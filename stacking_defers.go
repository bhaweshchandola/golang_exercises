package main

import "fmt"

// defer statement defers the executing of the function until the surrounding function returns
// i.e in case if it is added in print command until the main function returns it will not run
// stacking multiple defers will execute in LIFO order
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
