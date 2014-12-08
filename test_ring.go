package main

import (
    "fmt"
    "container/ring"
)

type A struct {
    n int
}

func main() {
    r := ring.New(5)



    n := r.Len()
    p := r
    for i := 0; i < n; i++ {
        val := i * 10
        p.Value = val
        fmt.Printf("set %d for i=%d\n", val, i)
        p = p.Next()
    }

    fmt.Printf("------\n")

    p = r
    for i := 0; i < n * 2; i++ {
        fmt.Printf("[%d]=%d\n", i, p.Value)
        p = p.Next()
    }
}