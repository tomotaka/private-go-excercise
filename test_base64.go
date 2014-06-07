package main

import (
    "fmt"
    "encoding/base64"
)

func main() {
    message := "hello world Base64!"
    messageBytes := []byte(message)

    b64encoded := base64.StdEncoding.EncodeToString(messageBytes)

    decoded, _ := base64.StdEncoding.DecodeString(b64encoded)

    fmt.Printf("input: %s\n", message)
    fmt.Printf("base64-encoded: %s\n", b64encoded)
    fmt.Printf("decode: %s\n", decoded)
}