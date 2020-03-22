package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "This string contains repeated elements like string contains and others"
	words := strings.Fields(x)
	// fmt.Println(words)
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		if _, ok := m[words[i]]; ok {
			m[words[i]] += 1
		} else {
			m[words[i]] = 1
		}
	}
	fmt.Println(m)
}
