package main

import "fmt"


func peakIndexInMountainArray(A []int) int {
    var tmpMaxIndex int = 0
    for i,v := range A[1:] {
        if v > A[tmpMaxIndex] {
            tmpMaxIndex = i+1
        }
    }
    return tmpMaxIndex
}

func main() {
    A := []int{1,2,1,13,4}
    fmt.Println(peakIndexInMountainArray(A))
}
