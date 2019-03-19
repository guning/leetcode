package main

import "fmt"


func minDeletionSize(A []string) int {
    num := 0
    for i := 0; i < len(A[0]); i++ {
        for j := 0; j < len(A) - 1; j++ {
            fmt.Println([]rune(A[j])[i], []rune(A[j+1])[i], i, j, num)
            if []rune(A[j])[i] > []rune(A[j+1])[i] {
                num++
                break
            }
        }
    }
    return num
}

func main() {
    A := []string{"cba","daf","ghi"}
    fmt.Println(minDeletionSize(A))
}
