package main

import (
    "fmt"
    "crypto/md5"
    "encoding/hex"
)

func main() {
    msg := "hello world"
    msgBytes := []byte(msg)

    md5hasher := md5.New()
    md5hasher.Write(msgBytes)

    hash := md5hasher.Sum(nil)

    fmt.Printf("msg: %s\n", msg)
    fmt.Printf("hash: %s\n", hex.EncodeToString(hash))
}