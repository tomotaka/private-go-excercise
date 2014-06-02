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

func prime_checker(feed_ch chan []int, result_ch chan []int, stop_sigch chan bool) {
    for {
        numbers := <-feed_ch
        if numbers == nil {
            break
        }

        length := len(numbers)
        tmp := make([]int, length)
        tmp_idx := 0
        for i := 0; i < length; i++ {
            n := numbers[i]
            if is_prime(n) {
                tmp[tmp_idx] = n
                tmp_idx++
            }
        }

        // copy
        result := make([]int, tmp_idx)
        for i := 0; i < tmp_idx; i++ {
            result[i] = tmp[i]
        }

        result_ch <- result
    } 
    stop_sigch <- true
    fmt.Println("checker dies")
}

func main() {
    max := 100000
    chunk_len := 100
    goroutine_n := 16
    feed_ch := make(chan []int, 100)
    prime_ch := make(chan []int, 100)
    checker_stop_sigch := make(chan bool)

    for i := 0; i < goroutine_n; i++ {
        go prime_checker(feed_ch, prime_ch, checker_stop_sigch)
        fmt.Printf("checker %d started!\n", i+1)
    }

    go func() {
        var chunk_idx int = 0
        var chunk []int = make([]int, chunk_len)
        for i := 2; i <= max; i++ {
            chunk[chunk_idx] = i
            chunk_idx++
            if chunk_idx == chunk_len {
                feed_ch <- chunk
                chunk = make([]int, chunk_len)
                chunk_idx = 0
            }
        }

        if 0 < chunk_idx {
            tmp := make([]int, chunk_idx)
            for i := 0; i < chunk_idx; i++ {
                tmp[i] = chunk[i]
            }
            feed_ch <- chunk
        }
        fmt.Println("gave numbers!")

        for i := 0; i < goroutine_n; i++ {
            feed_ch <- nil
            <-checker_stop_sigch
        }
        fmt.Println("killed all checkers!")
        prime_ch <- nil
    }()

    prime_count := 0
    for {
        prime_numbers := <-prime_ch
        if prime_numbers == nil {
            break
        }

        prime_count += len(prime_numbers)
        // fmt.Printf("add %d\n", len(prime_numbers))
    }



    fmt.Printf("total prime numbers = %d\n", prime_count)
    
}