package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/**
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

func _walk(t *tree.Tree, ch chan int) {
	if (t != nil) {
		if t.Left != nil {
			_walk(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			_walk(t.Right, ch)
		}
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	same := true
	for {
		v1, ok1 := <-c1 // receive from c1
		v2, ok2 := <-c2 // receive from c2
		if (ok1 && ok2) {
			if (v1 != v2) {
				same = false
				break;
			}
		} else {
			if (ok1 != ok2) {
				same = false;
			}
			break;
		}
	}
	return same
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
