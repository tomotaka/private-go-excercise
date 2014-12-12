package main

import (
    "fmt"

)


func MergeSort(data []int) []int {
    length := len(data)

    if length == 1 {
        return data
    }

    var a, b []int

    if length % 2 == 1 {
        a = make([]int, length / 2 + 1)
        b = make([]int, length / 2)
    } else {
        a = make([]int, length / 2)
        b = make([]int, length / 2)
    }

    aLen := len(a)
    bLen := len(b)

    for i := 0; i < aLen; i++ {
        a[i] = data[i]
    }

    j := 0
    for i := aLen; i < length; i++ {
        b[j] = data[i]
        j++
    }

    a = MergeSort(a)
    b = MergeSort(b)

    ia, ib, ir := 0, 0, 0

    ret := make([]int, length)
    for ia < aLen && ib < bLen {
        va := a[ia]
        vb := b[ib]

        if va < vb {
            ret[ir] = va
            ia++
        } else {
            ret[ir] = vb
            ib++
        }
        ir++
    }
    for ; ia < aLen; ia++ {
        ret[ir] = a[ia]
        ir++
    }
    for ; ib < bLen; ib++ {
        ret[ir] = b[ib]
        ir++
    }

    return ret
}

func print(data []int) {
    fmt.Printf("--------\n")
    length := len(data)
    for i := 0; i < length; i++ {
        fmt.Printf("%d\n", data[i])
    }
}

func main() {
    hoge := []int{ 1, 4, 3, 10, 2, 3, 30, 31, 498, 62, 22 }
    print(hoge)

    sorted := MergeSort(hoge)
    print(sorted)
}