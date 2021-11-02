package main

import (
	"fmt"
	"sync"
)

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"Title: The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Title: Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Title: Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Title: Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

// Fetcher ...
type Fetcher interface {
	Fetch(urlstring string) (urllist []string, err error)
}

func (f fakeFetcher) Fetch(urlstring string) ([]string, error) {
	if res, ok := f[urlstring]; ok {
		fmt.Printf("found: %s\n", urlstring)
		return res.urls, nil
	}
	fmt.Printf("missing: %s\n", urlstring)
	return nil, fmt.Errorf("not found: %s", urlstring)
}

// ###### Serial crawler ######
func Serial(url string, fetcher Fetcher, fetched map[string]bool) {
	if fetched[url] {
		return
	}
	fetched[url] = true
	urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	for _, u := range urls {
		Serial(u, fetcher, fetched)
	}
	return
}

// ###### Concurrent crawler with shared state and Mutex ######

func makeState() *fetchState {
	f := &fetchState{}
	f.fetched = make(map[string]bool)
	return f
}

type fetchState struct {
	mu      sync.Mutex
	fetched map[string]bool
}

func ConcurrentMutex(url string, fetcher Fetcher, f *fetchState) {
	f.mu.Lock()
	if f.fetched[url] {
		f.mu.Unlock()
		return
	}
	f.fetched[url] = true
	f.mu.Unlock()

	urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	var done sync.WaitGroup
	for _, u := range urls {
		done.Add(1)
		go func(u string) {
			defer done.Done()
			ConcurrentMutex(u, fetcher, f)
		}(u)
	}
	done.Wait()
	return
}

// ###### Concurrent crawler with channels ######

func worker(url string, ch chan []string, fetcher Fetcher) {
	urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- []string{}
	} else {
		ch <- urls
	}
}

func master(ch chan []string, fetcher Fetcher) {
	n := 1
	fetched := make(map[string]bool)
	for urls := range ch {
		for _, u := range urls {
			if !fetched[u] {
				fetched[u] = true
				n++
				go worker(u, ch, fetcher)
			}
		}
		n--
		if n == 0 {
			break
		}
	}
}

// ConcurrentChannel ...
func ConcurrentChannel(url string, fetcher Fetcher) {
	ch := make(chan []string)
	go func() {
		ch <- []string{url}
	}()
	master(ch, fetcher)
}

func main() {
	fmt.Printf("=== ConcurrentChannel ===\n")
	ConcurrentChannel("http://golang.org/", fetcher)
}
