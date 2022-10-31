package main

import (
	"fmt"
)

// 第二种写法: 左闭右开
func b_s(nums []int, target int) int {
	var left, right int
	right = len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := b_s(nums, 3)
	if res != -1 {
		fmt.Printf("找到了，目标索引为:%d\n", res)
	} else {
		fmt.Println("没找到目标值")
	}
}
