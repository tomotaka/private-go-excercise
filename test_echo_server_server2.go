package main

import (
    "log"
    "net"
    "bufio"
)

func same(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }

    l := len(a)
    for i := 0; i < l; i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func main() {
    l, err := net.Listen("tcp", ":2000")
    if err != nil {
        log.Fatal(err)
    }

    defer l.Close()

    for {
        conn, err := l.Accept()

        if err != nil {
            log.Fatal(err)
        }

        go func(c net.Conn){
            reader := bufio.NewReader(c)
            writer := bufio.NewWriter(c)

            for {
                line, _, err := reader.ReadLine()

                if err != nil {
                    log.Fatal(err)
                    break
                } else if same(line, []byte("exit")) {
                    writer.Write([]byte("bye\n"))
                    writer.Flush()
                    break
                }

                writer.Write(line)
                writer.Write([]byte{'\n'})  // ReadLine() does not return linebreak code
                writer.Flush()
            }

            c.Close()
        }(conn)
    }
}
