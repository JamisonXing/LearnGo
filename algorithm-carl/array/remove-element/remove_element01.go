package main

import "fmt"

// 解法二：双指针法，通过快慢指针在一个for循环中完成两个for循环的操作
func removeElement01(nums []int, val int) int {
	var slowIndex int
	for fastIndex, v := range nums {
		if v != val { //这里不是==，仔细想下
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}

	return slowIndex
}

func main() {
	nums := []int{1, 2, 4, 4, 4, 5}
	fmt.Printf("删除指定元素后，切片长度为:%d\n", removeElement01(nums, 4))
}
