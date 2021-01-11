package main

import (
    "flag"
    "fmt"
)

func main() {
    var (
	num = flag.Int("count", 0, "count")
	fname = flag.String("list")
    )
    flag.Parse()
    fmt.Println(*num)
    fmt.Println(*fname)
}

