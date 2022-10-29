package main

import (
	"map-memory-leak/utils"
	"runtime"
)

//polarisxu map竟然会内存泄漏
//map 总是可以在内存中增长；它从不收缩。
//因此，如果它导致一些内存问题，你可以尝
//试不同的选项，例如强制 Go 重新创建 map 或使用指针

/*每个m值都是一个 128 字节的数组。我们将执行以下操作：
分配一个空的 map。
添加 100 万个元素。
删除所有元素，并运行垃圾收集（GC）。*/

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

	runtime.GC() //triggers a manual gc
	utils.PrintAlloc()
	runtime.KeepAlive(m) //keep a reference to m so that the map isn't collected
}

/*
result:
105 MB <-- m 被分配后 //为什么不是0因为调用printAlloc()函数造成的
472503 MB <-- 我们增加 100 万个元素后
300440 MB <-- 我们删除 100 万个元素后
*/

/*
原因是 map 中的存储桶数无法缩减。因此，从 map 中删除元素不会影响现有存储桶的数量；
它只是将桶中的槽清零。map 只能增长并有更多的桶；它从不收缩。
*/

/*
解决方案:
一种解决方案可以是定期重新创建当前 map 的副本。
例如，每小时，我们可以构建一个新 map，复制所有
元素，然后发布上一个元素。此选项的主要缺点是，在
复制之后直到下一次垃圾回收之前，我们可能会在短时
间内消耗两倍于当前内存的内存。

另一个解决方案是更改 map 类型以存储数组指针：
map[int]*[128]byte。这并不能解决我们将拥有大量
桶的事实；然而，每个 bucket 条目将为该值保留指针
的大小，而不是 128 字节（64 位系统为 8 字节，32 位系统为 4 字节）。

见测试test00,test01
*/
