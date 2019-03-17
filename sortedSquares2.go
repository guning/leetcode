package main

import (
    "fmt"
    "sort"
)

type sortInt []int

func (B sortInt) Len() int {
    return len(B)
}

func (B sortInt) Less(i, j int) bool {
    return B[i] <= B[j]
}

func (B sortInt) Swap(i, j int) {
    B[i], B[j] = B[j], B[i]
}

func sortedSquares(A []int) []int {
    for i, v := range A {
        A[i] = v * v
    }
    sort.Sort(sortInt(A))
    return A
}


func main() {
	A := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(A)
	fmt.Println(res)
}
