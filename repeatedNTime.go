package main

func repeatedNTime(A []int) int {
    var repeatTimes = make(map[int]int)
    N := len(A) / 2
    for _, v := range A {
        repeatTimes[v] += 1
        if repeatTimes[v] == N {
            return v
        }
    }
    return v
}
