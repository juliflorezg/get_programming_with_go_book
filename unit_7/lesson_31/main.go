package main

import (
	"sync"
	"time"
)

// visited tracks web pages that have been visited
// its methods may be used concurrently from multiple goroutines
type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()

	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

var mu sync.Mutex

func main() {
	// mutex -> mutual exclusion, it serves to tell other goroutines to not do anything on a shared value for a while (locks & unlocks)
	// mu.Lock()
	// defer mu.Unlock()

	// visitedLinks := Visited{}
	// visitedLinks.visited = make(map[string]int)

	// webSites := []string{"google", "facebook", "instagram", "amazon", "youtube", "google", "facebook", "instagram", "amazon", "youtube", "instagram", "facebook", "google", "youtube"}

	// for _, l := range webSites {
	// 	randomTime := rand.Intn(len(webSites)/2)
	// 	time.Sleep(time.Duration(randomTime) * time.Second)

	// 	visitedLinks.VisitLink(l)
	// }

	// fmt.Println(visitedLinks.visited)

	// go worker()

	// time.Sleep(10 * time.Second)

	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Right()
	r.Start()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
}
