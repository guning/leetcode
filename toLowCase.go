package main

import (
    "fmt"
)

func toLowCase(str string) string {
    res := []rune(str)
    for i, c := range res {
        if c >= 65  && c <= 90 {
            res[i] = c + 32
        }
    }
    return string(res)
}


func main() {
    str := toLowCase("sSADWQEFEFEsssdSADSADASadad")
    fmt.Printf(str)
}
