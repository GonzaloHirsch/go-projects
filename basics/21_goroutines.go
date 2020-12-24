package main

import (
	"fmt"
	"time"
)

// Goroutines
// Goroutines are lightweight threads, they run in the same memory address and share memory, they need synchronization
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Channels
// Channels are typed data used to send and receive information
// The channel operator is <- (data flows in the direction of the arrow)
// They must be created before use, and send/receive blocks when the other side is ready
/*
	ch <- v    	// Send v to channel ch.
	v := <-ch  	// Receive from ch, and
				// assign value to v.
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// Range and close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	// Goroutines
	// Goroutines are specified with go keyword
	go say("world")
	say("hello")

	// Channels
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // Creating the channel c
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// Buffered channels
	// Channels can be buffered by providing the length in the make statement
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Range and close
	// Channels can be closed (it's better if the sender closes them)
	// One can test if the channel is closed via: v, ok := <-ch
	// One can iterate a channel until it is closed
	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}
}
