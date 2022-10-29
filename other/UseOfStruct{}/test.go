package main

import (
	"fmt"
	"unsafe"
)

// 空结构体，多用来做控制，传递信息
func main() {
	var s struct{}
	var inter interface{}
	var x [10000]struct{}
	var k = make([]struct{}, 100000)
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(k))
	fmt.Print(unsafe.Sizeof(inter))
}
