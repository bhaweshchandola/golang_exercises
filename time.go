package main

import (
	"fmt"
	"time"
)

func greet() string {
	return "hello world current time is: " + time.Now().String()
}

func main() {
	fmt.Println(greet())
}
