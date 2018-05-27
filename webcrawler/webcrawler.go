package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Fetch struct {
	url   string
	depth int
}

type Result struct {
	url  string
	body string
}

func loadUrls(out chan Fetch, fetches ...Fetch) {
	for _, fetch := range fetches {
		out <- fetch
	}
}

func filterDups(in chan Fetch, out chan Fetch) {
	known := make(map[string]bool)

	for fetch := range in {
		fmt.Printf("filter(): fetch = %v\n", fetch)
		switch {
		case fetch.depth <= 0:
			fmt.Println("  > depth achieved; pruning.")
		case known[fetch.url]:
			fmt.Printf("  > already fetched %v; pruning.\n", fetch.url)
		default:
			known[fetch.url] = true
			out <- fetch
		}
	}
	fmt.Println("filter(): DONE!")
}

func doFetch(in chan Fetch, out chan Result, urlsToFetch chan Fetch, fetcher Fetcher) {
	for fetch := range in {
		fmt.Printf("doFetch(): fetch = %v\n", fetch)
		body, urls, err := fetcher.Fetch(fetch.url)
		if err != nil {
			fmt.Printf("WARN: fetch of %v failed: \"%v\"\n", fetch.url, err)
			continue
		}
		out <- Result{fetch.url, body}
		for _, url := range urls {
			urlsToFetch <- Fetch{url, fetch.depth - 1}
		}
	}
	fmt.Println("doFetch(): DONE!")
	close(out)
}

func main() {
	urlsToProcess := make(chan Fetch, 20)
	fetches := make(chan Fetch)
	results := make(chan Result)

	go loadUrls(urlsToProcess, Fetch{"https://golang.org/", 4})
	go filterDups(urlsToProcess, fetches)
	go doFetch(fetches, results, urlsToProcess, fetcher)

Loop:
	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(1 * time.Second):
			fmt.Println("Done.")
			break Loop
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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
}
