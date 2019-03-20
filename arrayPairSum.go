package main

import "fmt"

func arrayPairSum(nums []int) int {
    sortedNums := quickSort(nums)
    sum := 0
    for i:=0; i<len(nums); i=i+2 {
        sum += sortedNums[i]
    }
    return sum
}

func quickSort(nums []int) []int {
    if len(nums) <=1 {
        return nums
    }
    i := 0
    j := len(nums) - 1
    for i!=j {
        for (nums[j] >= nums[0] && i < j) {
            j--
        }
        for (nums[i] <= nums[0] && i < j) {
            i++
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    nums[0], nums[i] = nums[i], nums[0]
    left := quickSort(nums[:i])
    right := quickSort(nums[i+1:])
    newNums := append(left, nums[i])
    newNums = append(newNums, right...)
    return newNums
}

func main() {
    nums := []int{1,4,2,3}
    fmt.Println(arrayPairSum(nums))
}
