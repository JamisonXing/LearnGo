package main

import "fmt"

// 测试go二维数组在空间上是不是连续的
// Test whether the go's twe-dimensional array is spatially continuous.
func main() {
	var m [2][2]int
	m = [2][2]int{
		{1, 1},
		{2, 2},
	}
	fmt.Print(&m[0][0], &m[0][1])
	fmt.Println()
	fmt.Print(&m[1][0], &m[1][1])
}

/*
result:
0x1400001c020 0x1400001c028
0x1400001c030 0x1400001c038
内存上是连续的
*/
