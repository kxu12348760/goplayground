package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n , lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim;
}

func Sqrt(x float64) float64 {
	guess := 2.0
	for i := 0; i < 10; i++ {
		guess -= (guess * guess - x) / (2 * guess)
	}
	return guess
}

func Sqrt2(x float64) float64 {
	guess := 23.0
	threshold := 0.0001
	i := 0
	close_enough := false
	for !close_enough {
		old_guess := guess
		guess -= (guess * guess - x) / (2 * guess)
		i += 1
		diff := math.Abs(guess - old_guess)
		close_enough = (diff < threshold)
	}
	fmt.Printf("%d iterations to guess square root of %g\n", i, x)
	return guess
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(sqrt(2), sqrt(3), sqrt(4), sqrt(5), sqrt(6))
	fmt.Println(Sqrt(2), Sqrt(3), Sqrt(4), Sqrt(5), Sqrt(6))
	fmt.Println(Sqrt2(2), Sqrt2(3), Sqrt2(4), Sqrt2(5), Sqrt2(6))
}
