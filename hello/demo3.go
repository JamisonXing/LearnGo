package main

import "fmt"

func main() {
	crr := []int{3, 5, 6}
	//subSlice()
	updateSlice(crr)
	fmt.Print(crr[0]) //修改成功
}

func subSlice() {
	arr := make([]int, 3, 5)
	crr := arr[0:2]
	crr[0] = 3
	fmt.Println(arr[0]) //arr也被修改

	crr = append(crr, 5, 6, 5, 6) //超过cap，内存分离
	crr[1] = 4
	fmt.Println(arr) //arr[1]并没有被改变
}

func updateSlice(crr []int) {
	crr[0] = 100
}
