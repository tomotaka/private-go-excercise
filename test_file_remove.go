package main

import (
    "os"
    "fmt"
    "log"
)

func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func main() {
    fname := "/tmp/hoge.txt"

    if !Exists(fname) {
        fmt.Printf("file %s does not exist, cannot remove it!\n", fname)
        return
    }

    err := os.Remove(fname)
    if err != nil {
        log.Fatal(err)
        return
    }

    fmt.Printf("removed %s\n", fname)
}