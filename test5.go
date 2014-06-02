package main

import (
    "fmt"
)

func is_prime(n int) bool {
    for i := 2; i < n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    count := 0
    max := 100000

    for i := 2; i <= max; i++ {
        if is_prime(i) {
            count++
        }
    }
    fmt.Printf("1-%d: %d prime numbers\n", max, count)
    
}