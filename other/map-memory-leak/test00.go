package main

import (
	"map-memory-leak/utils"
	"runtime"
)

// 用指针引用value,减少内存使用，并不能解决我们存在大量bucket的事实
func main() {
	n := 1_000_000
	k := make(map[int]*[128]byte)
	utils.PrintAlloc()

	for i := 0; i < n; i++ {
		k[i] = &[128]byte{}
	}
	utils.PrintAlloc()

	for i := 0; i < n; i++ {
		delete(k, i)
	}

	runtime.GC() //triggers a manual gc
	utils.PrintAlloc()
	runtime.KeepAlive(k) //keep a reference to m so that the map isn't collected
}

/*
result:
105 MB <-- m 被分配后 //为什么不是0因为调用printAlloc()函数造成的
186581 MB <-- 我们增加 100 万个元素后
39283 MB <-- 我们删除 100 万个元素后
*/
