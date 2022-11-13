package main

//https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/submissions/

func missingNumber(nums []int) int {
	hash := map[int]bool{}
	for _, v := range nums {
		hash[v] = true
	}

	for i := 0; ; i++ {
		if !hash[i] {
			return i
		}
	}
}
