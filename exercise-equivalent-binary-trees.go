package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walker(t *tree.Tree, ch chan int) {
	if t != nil {
		walker(t.Left, ch)
		ch <- t.Value
		walker(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walker(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// 終了合図
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		// 異なる葉を持つ合図
		if v1 != v2 {
			return false
		}
	}
}

// https://go-tour-jp.appspot.com/concurrency/8
func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
