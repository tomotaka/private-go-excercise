package main

import (
    "fmt"
)

func main() {
    var sum func(int, int) int = func(a, b int) int {
        return a + b
    }


    fmt.Printf("1 + 2 = %d\n", sum(1, 2));
    
}
