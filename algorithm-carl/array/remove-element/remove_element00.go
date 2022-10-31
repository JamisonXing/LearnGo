package main

import "fmt"

/*
ltc 27
book p40
给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
//数组空间是连续的，所谓的移除只是覆盖而已

//解法一:暴力解法，一个for遍历，一个for更新

func removeElement(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for k := i + 1; k < len(nums); k++ {
				nums[k-1] = nums[k]
			}
			i-- //这里我想的有点久，容易忘
			size--
		}
	}
	return size
}

func main() {
	nums := []int{1, 2, 4, 4, 4, 5}
	fmt.Printf("删除指定元素后，切片长度为:%d\n", removeElement(nums, 4))
}
