package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("lis.go>")
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)

		fmt.Print("lis.go>")
	}
}
