package main

import (
	"fmt"
	"unsafe"
)

type kong struct {
}

type kong_ struct {
	kong
	name string
}

func emptyStruct() {
	a := kong{}
	b := 2
	c := kong_{}
	d := kong{}
	fmt.Printf("空结构体a的大小：%d,空结构体的地址：%p\n",
		unsafe.Sizeof(a), &a)
	fmt.Printf("变量b的地址：%p\n", &b)
	fmt.Printf("非独立结构体c的大小：%d,非独立空结构体的地址：%p\n",
		unsafe.Sizeof(c), &c.kong)
	fmt.Printf("空结构体d的大小：%d,空结构体d的地址：%p\n",
		unsafe.Sizeof(d), &d)

	/*
		结果：
		空结构体a的大小：0,空结构体的地址：0xdc9540
		变量b的地址：0xc00001c098
		嵌套结构体c的大小：16,嵌套后空结构体的地址：0xc000062250
		空结构体d的大小：0,空结构体d的地址：0xdc9540
		结论：
		所有独立的空结构体都具有相同的地址，这个地址叫zeroBase,且该空结构体的大小为0
	*/
}

func implHashSet() { //hashset是只有k无v的数据结构
	//使用map实现hashset
	m := map[string]struct{}{}
	m["a"] = struct{}{}
	fmt.Println(m)
	/*
		结果：map[a:{}]
	*/
}

func implPureMes() {
	//传递空信息，在不需要实际意义的信息时可以优化内存的占用
	pureMess := make(chan struct{})
	fmt.Println(pureMess)
}

func stringDemo() {
	fmt.Println(unsafe.Sizeof("邢某某"))
	fmt.Println(unsafe.Sizeof("邢某某的blog"))
	/*
		结果：
		16
		16
	*/
}

func utf_8Test() {
	str := "英文名为Jamison"
	fmt.Println(len(str)) //19
	//四个汉字每个汉字三个字节，其余英文每个一字节
	//所以含有一个字节以上的字符不能使用传统for循环遍历，应该使用for range
	for _, s := range str {
		fmt.Printf("%c", s) //result:英文名为Jamison
	}
	fmt.Println()
	//错误做法
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) //result:è±æåä¸ºJamison
	}
}

func qieFenStr() {
	//取前三个汉字
	str := "邢某某的blog"
	str = string([]rune(str)[:3])
	fmt.Println(str) //result:邢某某
}

func initializeMap() {
	//make初始化
	m := make(map[string]string, 10)
	m["name"] = "xxj"
	m["age"] = "12"
	fmt.Println(m)
	//字面量初始化
	k := map[string]int{
		"1": 2,
		"2": 2,
		"3": 2,
	}
	fmt.Println(k)
	//删除
	delete(m, "name")
	fmt.Println(m)
	//打印
	for mes := range m {
		fmt.Println("年龄是：", m[mes])
	}

}

func main() { // 底层原理学习
	/*	emptyStruct()
		implHashSet()
		implPureMes()
		stringDemo()
		utf_8Test()*/
	qieFenStr()
	initializeMap()
}
