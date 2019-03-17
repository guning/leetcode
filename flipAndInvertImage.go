package main

func flipAndInvertImage(A [][]int) [][]int {
    for i,row := range A {
        var newRow []int
        for j:=len(row)-1; j>=0; j-- {
            newRow = append(newRow, row[j]^1)
        }
        A[i] = newRow
    }
    return A
}

func main() {
    A := [][]int{{1,0,0}, {1,0,1}, {0,0,1}}
    flipAndInvertImage(A)
}
