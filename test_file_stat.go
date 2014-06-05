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
    if len(os.Args) < 2 {
        fmt.Println("need arg file!")
        return
    }

    fname := os.Args[1]
    if !Exists(fname) {
        fmt.Printf("file %s does not exist!\n", fname)
        return
    }

    stat, err := os.Stat(fname)
    if err != nil {
        log.Fatal(err)
        return
    }

    fmt.Printf("file: %s\n", fname)
    fmt.Printf("name=%s\n", stat.Name())
    fmt.Printf("size=%d\n", stat.Size())
    fmt.Printf("mode=%s\n", stat.Mode().String())
    fmt.Printf("isDir=%v\n", stat.IsDir())
}
