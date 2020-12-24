package main

/*
Taken from: https://tour.golang.org/concurrency/10

In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
*/

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (fetcher *fakeFetcher) Crawl(url string, depth int) {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	// Checking the cache
	fetcher.mu.Lock()
	isUrlInCache := fetcher.cache[url]
	fetcher.mu.Unlock()
	// If not in cache, search for it and the recursive ones
	if !isUrlInCache {
		body, urls, err := fetcher.Fetch(url)
		// Add to cache
		fetcher.mu.Lock()
		fetcher.cache[url] = true
		fetcher.mu.Unlock()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		// Create channel to wait for kid
		done := make(chan bool)
		// Go to response urls
		for _, u := range urls {
			go func(url string) {
				fetcher.Crawl(url, depth-1)
				done <- true
			}(u)
		}
		// Waiting for responses
		for range urls {
			<-done
		}
	} else {
		fmt.Printf("Already looked for: %s\n", url)
	}
	return
}

func main() {
	fetcher.Crawl("https://golang.org/", 4)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
	data  map[string]*fakeResult
	mu    sync.Mutex
	cache map[string]bool
}

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f.data[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{data: map[string]*fakeResult{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
},
	cache: make(map[string]bool),
}
