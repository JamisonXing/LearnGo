package main

import (
	"fmt"
	"strings"
	"unicode"
)

func useOfBackticks() {
	fmt.Print(`hello
自动换行`)
	/*
	   hello
	   自动换行
	*/
}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && (!unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left]))) {
			left++
		}
		for left < right && (!unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right]))) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}

	}
	return true
}

func xiaoXie(s string) {
	res := strings.ToLower(s)
	fmt.Print(res)
}

func isLetter(s string) {
	print(unicode.IsLetter(rune(s[0])))
}

// 验证原生map不能进行并发操作
func rWOfMap() {
	m := make(map[int]int)

	//不断地读
	go func() {
		for {
			_ = m[1]
		}
	}()

	//不断的写
	go func() {
		for {
			m[2] = 2
		}
	}()

	//阻塞
	select {}
}

func main() {
	//useOfBackticks()
	//isPalindrome("ab2a")
	//isLetter("2")

	rWOfMap()
}
