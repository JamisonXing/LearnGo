package day3

//https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

func replaceSpace(s string) string {
	b := []byte(s)
	res := make([]byte, 0)

	for _, v := range b {
		if v == ' ' {
			res = append(res, []byte("%20")...) //可变参数
		} else {
			res = append(res, v)
		}
	}

	return string(res)
}
