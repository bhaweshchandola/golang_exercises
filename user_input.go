package main

import (
	"fmt"
	// "bufio"
	// "os"
)

func main(){
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter any number: ")
	// num, _ := reader.ReadString('\n')
	// if(1<num<10){
	// 	fmt.Print("Number is between 1 and 10")
	// }
	var num int
	fmt.Scanf("%d", &num)
	if 1<num && num<10 {
		fmt.Println("number between 1 and 10")
	}
}