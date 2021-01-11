package main

import "fmt"

func main() {
	var s = [] int {}

	s = append(s, 2)
	s = append(s, 3)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
