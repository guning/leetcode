package main

import "fmt"

func main() {
    arr := make([]int, 4, 4)
    arr[1] = 222

    fmt.Println("cap: %d len %d", cap(arr), len(arr))
}
