package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	var p *Person

	p = &Person{
		Name: "太郎",
		Age: 20,
	}
	fmt.Printf("p: %+v\n", p)
	fmt.Printf("the address in the variable p: %p\n", p)
}
