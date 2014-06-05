package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    fname := "/tmp/test_file_write2.go.txt"

    contents := "Hello world"
    contentsBytes := []byte(contents)
    contentsLen := len(contentsBytes)

    fi, err := os.OpenFile(fname, os.O_WRONLY | os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
        return
    }

    totalWroteBytes := 0
    for totalWroteBytes < contentsLen {
        wroteBytes, err := fi.Write(contentsBytes)
        if err != nil {
            log.Fatal(err)
            return
        }

        totalWroteBytes += wroteBytes
    }

    err = fi.Close()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("wrote content to file: %s\n", fname)
}