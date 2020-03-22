package main

import "fmt"

type Vertex struct {
	key, value float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["some"] = Vertex{
		10.1, 10.2,
	}
	fmt.Println(m)
	var x = map[string]Vertex{
		"first":  {10, 12},
		"second": {30, 50},
	}
	fmt.Println(x)
}
