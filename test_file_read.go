package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    fileName := "./hoge.txt"
    fi, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
        return
    }

    buffer := make([]byte, 100)
    readLen, err := fi.Read(buffer)
    if err != nil {
        log.Fatal(err)
        return
    }
    available := make([]byte, readLen)
    if copy(available, buffer) != readLen {
        fmt.Println("something wrong happened")

    }
    content := string(available)
    fmt.Printf(content)
}
