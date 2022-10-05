package main

import "fmt"

type car interface {
	Drive()
}

type truck struct {
	driver string
}

func (t truck) Drive() {
	fmt.Println(t.driver + " is driving")
}

func main() {
	t := truck{
		driver: "xxj",
	}
	t.Drive()
}
