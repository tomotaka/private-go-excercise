package main

import (
    "fmt"
)

func is_prime(n int64) bool {
    var i int64;
    for i = 2; i < n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var count int64 = 0
    var max int64 = 100000

    var i int64

    for i = 2; i <= max; i++ {
        if is_prime(i) {
            count++
        }
    }
    fmt.Printf("1-%d: %d prime numbers\n", max, count)
    
}