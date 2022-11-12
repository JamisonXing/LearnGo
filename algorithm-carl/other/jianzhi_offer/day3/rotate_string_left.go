package day3

//https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=4j3h5s1

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b, 0, len(b)-1)
	reverse(b, 0, len(b)-n-1)
	reverse(b, len(b)-n, len(b)-1)
	return string(b)
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

/*
//耗时但简单的方法：
func reverseLeftWords(s string, n int) string {
    for i := 0; i < n; i++ {
        s += string(s[i])
    }
    s = s[n:]
    return s
}
*/
