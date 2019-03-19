package main

import "fmt"
import "math"

func selfDividingNumbers(left int, right int) []int {
    var res []int
    for i := left; i <= right; i++ {
        if (isSelfDividingNumber(i)) {
            res = append(res, i)
        }
    }
    return res
}

func isSelfDividingNumber(num int) bool {
    newNum := num
    for newNum > 0 {
        re := newNum % 10
        newNum = int(math.Floor(float64(newNum/10)))
        if re == 0 {
            return false
        } else {
            if num % re != 0 {
                return false
            }
        }
    }
    return true
}

func main() {
    res := selfDividingNumbers(1,22)
    fmt.Println(res)
}
