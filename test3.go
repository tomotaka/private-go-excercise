package main 

import (
    "fmt"
    "container/list"
)

type A struct {
    n int
}


func main() {
    my_list := list.New()
    my_list.PushBack(A{n: 100})
    my_list.PushBack(A{n: 200})

    for e := my_list.Front(); e != nil; e = e.Next() {
        // fmt.Println(e)
        fmt.Printf("n=%d!!!!\n", e.Value.(A).n)
    }

    
}
