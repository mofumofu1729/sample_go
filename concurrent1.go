package main

import (
	"fmt"
	"testing"
	"time"
)

func main() {
	result := testing.Benchmark(func(b *testing.B) { run() })
	fmt.Println(result.T)
}

func run() {
	fmt.Println("Start")
	process("A")
	process("B")
	process("C")
	fmt.Println("Finish!")
}

func process(name string) {
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}
