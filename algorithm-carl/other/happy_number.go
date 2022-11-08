package main

import "fmt"

func step(n int) int {
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

func main() {
	n := 1
	if isHappy(n) {
		fmt.Printf("n is a happy number%d\n", n)
	} else {
		fmt.Printf("n isn't a happy number%d\n", n)
	}
}
