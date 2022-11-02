package main

import (
	"fmt"
)

//https://leetcode.cn/problems/squares-of-a-sorted-array/

//数组其实是有序的， 只不过负数平方之后可能成为最大数了。
//
//那么数组平方的最大值就在数组的两端，不是最左边就是最右边，不可能是中间。
//
//此时可以考虑双指针法了，i指向起始位置，j指向终止位置。

func SquareAndSort01(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	res := make([]int, n)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			res[k] = lm
			i++
		} else {
			res[k] = rm
			j--
		}
		k--
	}
	return res
}

func main() {
	nums := []int{-20, -19, -14, -12, -7, -5, 2, 2, 4, 6, 6, 7, 8, 9, 10, 12, 17, 17, 18, 18}
	fmt.Println(SquareAndSort01(nums))
}
