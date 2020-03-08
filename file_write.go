package main

import "os"

func main(){
	file, err := os.Create("file.txt")

	if err!=nil{
		return
	}

	defer file.Close()

	var a = []string{"first", "second", "third"}
	for x:=0;x<len(a);x++{
		file.WriteString(a[x])
		file.WriteString("\n")
	}
}