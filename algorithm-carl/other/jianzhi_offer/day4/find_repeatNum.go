package main

func findRepeatNumber(nums []int) int {
	m := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = struct{}{}
	}
	return -1
}
