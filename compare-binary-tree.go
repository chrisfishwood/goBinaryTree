package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1c := make(chan int)
	t2c := make(chan int)

	go func() {
		Walk(t1, t1c)
		close(t1c)
	}()
	go func() {
		Walk(t2, t2c)
		close(t2c)
	}()

	for t1v := range t1c {
		t2v := <-t2c
		if t1v != t2v {
			return false
		}
	}
	return true
}

func main() {
	same := Same(tree.New(1), tree.New(1))
	fmt.Println("same (should be true):", same)
	same = Same(tree.New(1), tree.New(2))
	fmt.Println("same (should be false):", same)
}
