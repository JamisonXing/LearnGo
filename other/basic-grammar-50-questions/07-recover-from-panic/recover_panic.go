package main

import "fmt"

//如何从panic中恢复
/*
在一个 defer 延迟执行的函数中调用 recover ，它便能捕捉/中断 panic。
*/

/*// 错误方式
func main() {
	recover() //什么都不做
	panic("bad")
	recover() //不会执行
	fmt.Println("ok")
}*/

// 正确方式
func main() {
	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	panic("bad")

	/*	m := make(map[string]any)
		m["one"] = 1
		m["two"] = 2
		m["thr"] = 3
		for s := range m {
			fmt.Println(s)
		}*/

}

/*
9.简短变量的声明需要注意什么
-简短变量的声明只能在函数内部使用
-struct的变量字段不能使用:=来赋值
-不能使用简短声明的变量重复声明， :=左侧至少有一个新变量，才允许多变量的重复声明
*/

/*
10.range迭代map是有序的吗
无序的，go的runtime是有意打乱迭代顺序的，所以你得道德迭代结果可能不一样。但也不总会打乱，连续得到5个连续的结果是可能的。
*/
