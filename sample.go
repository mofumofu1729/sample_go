package main

import (
	"fmt"
	"flag"
	"os"
	"log"
	"bufio"
	"strings"
)

func main() {
	var (
		list_flag string
	)
	flag.StringVar(&list_flag, "list", "list_path", "path to list")

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(2)
	}

	f, err := os.Open(list_flag)
	if err != nil {
		log.Fatal(err)
		return
	}
	var lines []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	for i := range lines {
		log.Print(strings.ToUpper(lines[i]))
	}
}
