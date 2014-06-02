package main 

import (
    "fmt"
    // "container/list"
)

func is_prime(n int) bool {
    for i := 2; i < n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func prime_checker(feed_ch chan int, result_chan chan int, stop_signal chan bool) {
    checked := 0
    for {
        n := <-feed_ch
        if n == 0 {
            break
        }
        checked++
        if is_prime(n) {
            result_chan <- n
        }
    }
    fmt.Printf("checked %d numbers\n", checked)
    stop_signal <- true
}

func prime_dumper(result_chan chan int, finish_signal chan bool) {
    count := 0
    for {
        n := <-result_chan
        if n == 0 {
            break
        }
        // fmt.Printf("%d\n", n)
        count++
    }

    fmt.Printf("all prime numbers: %d\n", count)

    finish_signal <- true
}

func main() {
    max := 200000
    goroutine_n := 16

    checker_stop_signal := make(chan bool)

    finish_signal := make(chan bool)
    feed_num_ch := make(chan int, 500)
    prime_num_ch := make(chan int)


    // run multiple prime checkers
    for i := 0; i < goroutine_n; i++ {
        go prime_checker(feed_num_ch, prime_num_ch, checker_stop_signal)
        fmt.Printf("checker %d started!\n", i+1)
    }

    // dump prime numbers
    go prime_dumper(prime_num_ch, finish_signal)

    // check all numbers
    for i := 2; i <= max; i++ {
        feed_num_ch <- i
    }

    for i := 0; i < goroutine_n; i++ {
        feed_num_ch <- 0  // kill all prime_checker()
        <-checker_stop_signal
    }

    // let prime_dumper finish
    prime_num_ch <- 0


    <-finish_signal
    fmt.Println("completed")
}