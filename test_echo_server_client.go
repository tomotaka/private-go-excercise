package main

import (
    "fmt"
    "log"
    "net"
    "bufio"
)

func main() {
    r_addr, err := net.ResolveTCPAddr("tcp", "localhost:2000")
    if err != nil {
        log.Fatal(err)
    }
    conn, err := net.DialTCP("tcp", nil, r_addr)

    msg := []byte("hello\n")

    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)

    writer.Write(msg)
    writer.Flush()

    line, _, err := reader.ReadLine()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", line)

    writer.Write([]byte("exit\n"))
    writer.Flush()

    err = conn.Close()
    if err != nil {
        log.Fatal(err)
    }
}