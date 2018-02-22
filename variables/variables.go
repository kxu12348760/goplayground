package main

import "fmt"

var c, python, java = true, false, "no!"

var i, j int = 1, 2

func main() {
	k := 3
	scale := "Scala"
	/* Will print default values for types */
	fmt.Println(i, j, c, python, java, k, scale)
}
