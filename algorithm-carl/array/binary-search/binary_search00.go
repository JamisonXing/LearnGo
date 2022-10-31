package main

import "fmt"

//ltc 704
//book p37
//二分法要用于有序且无重复元素的数组更佳
//二分法通过区间分类有两种写法,左闭右闭和左闭右开

// 第一种写法：左闭右闭
func bs(nums []int, target int) int {
	var left int
	var right int
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2 //防止溢出
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else {
			return mid
		}
	}

	//没找到目标值
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := bs(nums, 3)
	if res != -1 {
		fmt.Printf("找到了，目标索引为:%d\n", res)
	} else {
		fmt.Println("没找到目标值")
	}

}
