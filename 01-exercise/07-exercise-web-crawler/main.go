package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/html"
)

type url_msg struct {
	url       string
	cur_depth int
	max_depth int
}

var fetch_cnt int64 = 1

var fetched map[string]bool
var fetched_mu sync.RWMutex
//  Fetch add equivalent: ori = atomic.AddUint32(&avalue, 1) - 1
// runtime.GOMAXPROCS(4)
// var is_all_fetched sync.Mutex
// var c = sync.NewCond(&is_all_fetched)
// var fetch_ch = make(chan url_msg)

// Crawl uses findLinks to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int) {
	// TODO: Fetch URLs in parallel.
	defer atomic.AddInt64(&fetch_cnt, -1)
	if depth < 0 {
		// fmt.Println("Reached end!")
		return
	}
	urls, err := findLinks(url)
	if err != nil {
		// fmt.Println(err)
		return
	}

	fmt.Println("fetch_cnt: ", fetch_cnt)
	fmt.Printf("found: %s\n", url)
	// fetched_mu.RLock()

	// fetched_mu.RUnlock()
	for _, u := range urls {
		fetched_mu.RLock()
		isFetched := fetched[u]
		fetched_mu.RUnlock()
		// can't defer RUnlock here cuz the spawned go routines would try to acquire and then kena locked?  
		if !isFetched {
			fetched_mu.RLock()
			fetched[u] = true
			fetched_mu.RUnlock()
			atomic.AddInt64(&fetch_cnt, 1)		
			go Crawl(u, depth-1)
		} 
		
	}
}

func main() {
	fetched = make(map[string]bool)
	now := time.Now()
	fetched["http://andcloud.io"] = true
	go Crawl("http://andcloud.io", 2)
	for atomic.LoadInt64(&fetch_cnt) > 0 {
	}
	fmt.Println("time taken:", time.Since(now))
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
