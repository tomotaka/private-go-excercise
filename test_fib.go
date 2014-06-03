package main

import (
    "fmt"
)

func fib(n int) int {
    if n <= 2 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}

func main() {
    n := 29
    fib_result := fib(n)

    fmt.Printf("fib(%d)=%d\n", n, fib_result)
    
}