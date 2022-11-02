package main

import (
	"fmt"
	"sort"
)

//https://leetcode.cn/problems/squares-of-a-sorted-array/

// SquareAndSort 暴力法
func SquareAndSort(nums []int) []int {
	res := make([]int, len(nums), len(nums)) //这里不能写var res []int,因为这样没有分配空间，报错：panic: runtime error: index out of range [0] with length 0
	for i, v := range nums {
		res[i] = v * v
	}
	sort.Ints(res)
	return res
}

func main() {
	nums := []int{-20, -19, -14, -12, -7, -5, 2, 2, 4, 6, 6, 7, 8, 9, 10, 12, 17, 17, 18, 18}
	fmt.Println(SquareAndSort(nums))
}
