package main

import "fmt"

func sortedSquares(A []int) []int {
	var res []int
	iv := insertValue()
	for _, v := range A {
		res = iv(v * v)
	}
	return res
}

func insertValue() func(int) []int {
	var res []int
	return func(v int) []int {
		if len(res) != 0 {
			for i, ov := range res {
				if v >= ov {
                                        if i == len(res) - 1 {
                                            res = append(res, v)
                                        } else {
					    continue
                                        }
				} else {
					if i == 0 {
						res = append([]int{v}, res...)
					} else {
						r := make([]int, len(res)+1)
						copy(r[:i], res[:i])
						r[i] = v
						copy(r[i+1:], res[i:])
						res = r
					}
					break
				}
			}
		} else {
			res = append(res, v)
		}
		return res
	}
}

func main() {
	A := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(A)
	fmt.Println("sss", res)
}
