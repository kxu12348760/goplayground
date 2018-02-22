package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	// (%v, %T) == (value, Type)
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i interface{}) {
	// (%v, %T) == (value, Type)
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	describe(i)
	// i.M() // Error

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var i2 interface{} // empty interface
	describe2(i2)

	i2 = 42
	describe2(i2)

	i2 = "hello"
	describe2(i2)

	var i3 interface{} = "hello"

	s := i3.(string)
	fmt.Println(s)

	s, ok := i3.(string)
	fmt.Println(s, ok)

	// type assertion will not trigger a panic
	f, ok := i3.(float64)
	fmt.Println(f, ok)

	f = i3.(float64) // panic
	fmt.Println(f)
}
