package main

import "fmt"

func hammingDistance(x int , y int) int {
    z := x ^ y
    num := 0
    for z > 0 {
        z &= z-1
        num++
    }
    return num
}

func main() {
    z := hammingDistance(2, 4)
    fmt.Println(z)
}
