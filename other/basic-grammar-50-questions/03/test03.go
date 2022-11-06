package main

import "fmt"

//switch 中如何强制执行下一个代码块
/*switch 语句中的 case 代码块会默认带上 break，
但可以使用 fallthrough 来强制执行下一个 case 代码块。*/

func main() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ': //没有的fallthrough的情况下，空格会直接break，返回false,和其他语言不一样
			//fallthrough //返回true
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t')) // true
	fmt.Println(isSpace(' '))  // false
}
