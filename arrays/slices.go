package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Raul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	primes2 := []int{2, 3, 5, 7, 11, 13}
	printSlice("what", primes2)

	primes2 = primes2[:0]
	printSlice("hello", primes2)

	primes2 = primes2[:4]
	printSlice("world", primes2)

	primes2 = primes2[2:]
	printSlice("gogo", primes2)

	// primes2 = primes2[1:4]
	// fmt.Println(primes2)

	// primes2 = primes2[:2]
	// fmt.Println(primes2)

	// primes2 = primes2[1:]
	// fmt.Println(primes2)
	var snil []int
	fmt.Println(snil, len(snil), cap(snil))

	if snil == nil {
		fmt.Println("nil!")
	}

	a2 := make([]int, 5)
	printSlice("a", a2)

	b2 := make([]int, 0, 5)
	printSlice("b", b2)

	c2 := b2[:2]
	printSlice("c", c2)

	d2 := c2[2:5]
	printSlice("d", d2)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var appends []int
	printSlice("appends", appends)

	appends = append(appends, 0)
	printSlice("appends", appends)

	// The slice grows as needed.
	appends = append(appends, 1)
	printSlice("appends", appends)

	// We can add more than one element at a time.
	appends = append(appends, 2, 3, 4)
	printSlice("appends", appends)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
