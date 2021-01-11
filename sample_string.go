package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	//"log"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	ListPath string `long:"listpath" required:"true"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	fp, err := os.Open(opts.ListPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		signal, _ := strconv.Atoi(split[1])
		fmt.Printf("%s: %d\n", split[0], signal)
	}
}

