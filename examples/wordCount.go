package main

/*
Taken from: https://tour.golang.org/moretypes/23

Implement WordCount.
It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

// To make it available offline run "go get golang.org/x/tour/gotour"
import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	parts := strings.Fields(s)
	for _, v := range parts {
		_v, ok := m[v]
		if ok {
			m[v] = _v + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
