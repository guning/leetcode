package main

import "fmt"


func diStringMatch(S string) []int {
    min := 0
    max := len(S)
    var res []int
    for _, v := range []rune(S) {
        if v == 'D' {
            res = append(res, max)
            max--
        } else {
            res = append(res, min)
            min++
        }
    }
    res = append(res, max)
    return res
}

func main() {
    S := "DIDI"
    fmt.Println(diStringMatch(S))
}
