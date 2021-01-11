package main

import (
	"fmt"
	"time"

)
func main() {
	result := testing.Benchmark(func(b *testing.B) { run() })
	fmt.Println(result.T)
}

func run() {
	fmt.Println("Start!")

	ch := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()

	isFin := <-ch

	close(ch)

	fmt.Println(isFin)
	fmt.Println("Finih!")
}
