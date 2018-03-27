package main

import "fmt"
import "github.com/goplayground/codefights/construct_square"

func main() {
	var test []int = []int{4, 2, 1, 1, 1}
	var digits []int = []int{4, 2, 1}
	var idx int = 0
	for (idx < 20) {
		test = construct_square.GenNextPermutation(test, digits)
		fmt.Println(test)
		idx = idx + 1
	}
}
