package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		fmt.Println("t nil")
		return
	}

	fmt.Println("Walk called Left")
	Walk(t.Left, ch)
	fmt.Println("value:", t.Value)
	ch <- t.Value
	fmt.Println("Walk called Right")
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

	same := true
	ctr := 0
	for t1v := range t1c {
		ctr = ctr + 1
		if ctr == 100 {
			fmt.Println("aaaaahhhhhhhh")
			return false
		}
		//t1v := <-t1c
		t2v := <-t2c
		fmt.Println("comparing:", t1v, t2v)
		if t1v != t2v {
			same = false
		}
	}
	return same
}

func main() {
	same := Same(tree.New(1), tree.New(1))
	fmt.Println("same:", same)
}
