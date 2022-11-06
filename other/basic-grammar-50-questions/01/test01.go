package main

import (
	"fmt"
)

//访问map中的key需要注意什么
/*
当访问 map 中不存在的 key 时，Go 则会返回元素对应数据类型的零值，比如 nil、’’ 、false 和 0，
取值操作总有值返回，故不能通过取出来的值，来判断 key 是不是在 map 中。
检查 key 是否存在可以用 map 直接访问，检查返回的第二个参数即可。
*/

// IsKeyExist 错误的检测方法
func IsKeyExist(x map[string]string) {
	if v := x["one"]; v == "" {
		fmt.Println("isKeyExist:key is no entry") //key "two" 存不存在都会返回空字符串
	} else {
		fmt.Println("isKeyExist:key has entry")
	}
}

// IsKeyExist01 通过第二个参数访问
func IsKeyExist01(x map[string]string) {
	if _, ok := x["two"]; !ok {
		fmt.Println("IsKeyExist01: key is no entry")
	} else {
		fmt.Println("IsKeyExist01: key have entry")
	}
}

func main() {
	x := map[string]string{"one": "1"} //"two"键不存在
	IsKeyExist(x)
	IsKeyExist01(x)
}
