package demo

import (
	"fmt"
	"testing"
)

func TestDemo00(t *testing.T) {
	u1 := &User{}
	u1.Option(WithId(1))
	fmt.Printf("%+v\n", u1)

	u2 := &User{}
	u2.Option(WithId(1), WithName("jamison"))
	fmt.Printf("%+v\n", u2)
}
