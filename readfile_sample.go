package main

import (
	"bufio"
	"fmt"
	"flag"
	"os"
)

func main() {
	var fp *os.File
	var err error
	var (
		path = flag.String("filepath", "/path/to/file", "filepath")
	)

	flag.Parse()
	fmt.Println(*path)

 	fp, err = os.Open(*path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	
	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(">", line)
	}
}
