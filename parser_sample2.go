package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	var (
		s = flag.String("file_path", "/path/to/defalt_file", "string flag")
	)
	flag.Parse()
	fmt.Println(*s)
	fmt.Println(filepath.Base(*s))
}
