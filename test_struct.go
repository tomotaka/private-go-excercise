package main

import (
    "fmt"
)

type Rectangle struct {
    height int
    width int
}

func (r Rectangle) Area() int {
    return r.height * r.width;
}

func main() {
    r1 := Rectangle{5, 7};
    fmt.Printf("height=%d, width=%d\n", r1.height, r1.width)
    fmt.Printf("area=%d\n", r1.Area())
        
}
