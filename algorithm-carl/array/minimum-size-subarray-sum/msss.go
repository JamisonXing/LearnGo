package main

import "fmt"

//https://leetcode.cn/problems/minimum-size-subarray-sum/

// 滑动窗口实现
func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)
	sum := 0
	res := l + 1
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if subLen < res {
				res = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if res == l+1 {
		return 0
	} else {
		return res
	}
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(22, nums))
}
