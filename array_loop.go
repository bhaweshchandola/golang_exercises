package main

import "fmt"

func main() {
	var a = []int64{1,2,3,4,5,56,7,7,7}
	for x:=0;x<9;x++{
		if a[x]>5{
			fmt.Println("Number greater than 5")
		}
	}
}