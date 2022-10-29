package main

import (
	"map-memory-leak/utils"
	"runtime"
)

// 创建副本清理map消耗的内存
func main() {
	n := 1_000_000
	m := make(map[int][128]byte)
	utils.PrintAlloc()

	for i := 0; i < n; i++ {
		m[i] = [128]byte{}
	}
	utils.PrintAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	for i := 0; i < n/2; i++ { //加5000个元素
		m[i] = [128]byte{}
	}

	utils.PrintAlloc()
	k := m //创建副本

	runtime.GC()         //triggers a manual gc
	runtime.KeepAlive(k) //keep a reference to m so that the map isn't collected
	utils.PrintAlloc()
}

/*
result:
105 MB
472517 MB <----加一万elements之后
472461 MB <----源map一万个bucket,5000个element
300440 MB <----副本map五千个bucket，5000个element
*/
