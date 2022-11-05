package main

//使用nil值的sile, map会发生什么
/*允许对nil值的sile和map添加元素会造成运行时panic*/

func nilSlice() {
	var s []string
	s[0] = "jamison"
}

func nilMap() {
	var m map[string]string
	m["name"] = "jamison"
}

func main() {
	nilSlice()
	go nilMap()
}

/*
nilMap result:
panic: assignment to entry in nil map
nilSlice result:
panic: runtime error: index out of range [0] with length 0
*/
