package main

import (
	"fmt"
	"os"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	ListPath string `long:"listpath" required:"true"`
	GPU int `long:"gpu" default:"-1"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	fmt.Printf("ListPath: %s\n", opts.ListPath)
	fmt.Printf("GPU: %d\n", opts.GPU)
}

