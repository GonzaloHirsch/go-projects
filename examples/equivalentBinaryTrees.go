package main

/*
Taken from: https://tour.golang.org/concurrency/7 and https://tour.golang.org/concurrency/8

There can be many different binary trees with the same sequence of values stored in it. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

A function to check whether two binary trees store the same sequence is quite complex in most languages.
We'll use Go's concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
Continue description on next page.

1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here.
*/

// To make it available offline run "go get golang.org/x/tour/gotour"
import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// We want to run through from lowest to highest
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

// Recursive walking the tree
func walkRecursive(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkRecursive(t.Left, ch) // Going left
	}
	ch <- t.Value // Sending value
	if t.Right != nil {
		walkRecursive(t.Right, ch) // Going right
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Creating the channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	// Walking trees
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	// Checking if the same
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	s1 := Same(tree.New(1), tree.New(1))
	s2 := Same(tree.New(1), tree.New(2))
	fmt.Println("Test 1 was: %v", s1)
	fmt.Println("Test 2 was: %v", s2)
	/*
		t1, t2 := tree.New(1), tree.New(1)
		testWalk(t1)
		testWalk(t2)
	*/
}

func testWalk(t *tree.Tree) {
	// Creating the channels
	ch := make(chan int)
	// Walking trees
	go Walk(t, ch)
	// Checking if the same
	for {
		v1, ok1 := <-ch
		if !ok1 {
			break
		} else {
			fmt.Println(v1)
		}
	}
}

}
