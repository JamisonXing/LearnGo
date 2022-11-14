package day5

//https://www.bilibili.com/video/BV1Bf4y1F7SY/?spm_id_from=333.999.0.0&vd_source=1761deea5b8c69ccd1cd089a400b2c54

func minArray(numbers []int) int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		mid := l + (r-l)/2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}
	return numbers[r]
}
