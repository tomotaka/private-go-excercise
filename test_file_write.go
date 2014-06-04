package main

import (
    "io/ioutil"
    "log"
)

func main() {
    fname := "/tmp/gotest.txt"

    contents := "hello this is golang excercise, i am writing to /tmp/gotest.txt"
    contentsBytes := []byte(contents)

    err := ioutil.WriteFile(fname, contentsBytes, 0644)
    if err != nil {
        log.Fatal(err)
        return
    }
}