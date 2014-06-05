package main

import (
    "fmt"
    "sync"
)

func is_prime(n int) bool {
    for i := 2; i < n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func prime_checker(input chan int, output chan int, waiter *sync.WaitGroup) {
    defer func(){
        fmt.Println("done!")
        waiter.Done()
    }()

    for n := range input {
        if is_prime(n) {
            output <- n
        }
    }
}

func result_printer(result chan int, end_hook func()) {
    for prime_number := range result {
        fmt.Printf("%d\n", prime_number)
    }
    end_hook()
}

func feeder(start int, end int, feed chan int) {
    for i := start; i <= end; i++ {
        feed <- i
    }
    close(feed)
    fmt.Println("finished feeding")
}

func main() {
    max := 100
    goroutine_n := 4

    feed_ch := make(chan int)
    result_ch := make(chan int)
    waiter := new(sync.WaitGroup)

    go feeder(1, max, feed_ch)

    for i := 0; i < goroutine_n; i++ {
        waiter.Add(1)
        go prime_checker(feed_ch, result_ch, waiter)
    }

    end_hook_waiter := new(sync.WaitGroup)
    end_hook_waiter.Add(1)
    var end_hook func() = func() {
        defer end_hook_waiter.Done()
        fmt.Println("end!")
    }

    go result_printer(result_ch, end_hook)

    waiter.Wait()
    fmt.Println("@@@ closing...")
    close(result_ch)

    end_hook_waiter.Wait()
}