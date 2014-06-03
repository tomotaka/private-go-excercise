package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
)

func main() {
    fileName := "./hoge.txt"
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
        return
    }

    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
        return
    }

    contentString := string(data)
    fmt.Println(contentString)
}