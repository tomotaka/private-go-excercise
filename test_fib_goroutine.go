package main

import (
    "fmt"
)

func fib(n int, result_ch chan int) {
    if n <= 2 {
        result_ch <- 1
        return
    }
    child_ch := make(chan int, 2)
    go fib(n - 2, child_ch)
    go fib(n - 1, child_ch)
    result_ch <- (<-child_ch + <-child_ch)
}

func main() {
    n := 29

    fib_result_ch := make(chan int)
    go fib(n, fib_result_ch)

    fmt.Printf("fib(%d)=%d\n", n, <-fib_result_ch)
}
