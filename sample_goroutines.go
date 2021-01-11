package main

import (
    "fmt"
    "time"
)

func f(value string) {
    for i := 0; i < 3; i++ {
        fmt.Println(value)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go f("goroutineを使って実行")
    f("普通に実行")
    fmt.Println("done")
}
