package main

import "sort"

//https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

func search(nums []int, target int) int {
	return sort.SearchInts(nums, target+1) -
		sort.SearchInts(nums, target)
}
