package day5

func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, v := range cnt {
		if i == 1 {
			return s[v]
		}
	}
	return ' '
}
