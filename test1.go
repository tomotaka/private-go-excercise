package main

import (
    "fmt"
)

func main() {
    c := make(chan bool)
    x := func(channel chan bool) {
        fmt.Println("this is x")
        channel <- true
    }

    go x(c)
    fmt.Println("hello")
    <-c
}