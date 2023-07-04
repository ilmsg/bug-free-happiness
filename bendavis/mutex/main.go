package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type SafeCounter struct {
	mu     sync.Mutex
	NumMap map[string]int
}

func (s *SafeCounter) Add(num int) {
	// s.mu.Lock()
	// s.NumMap["key"] = num
	s.mu.Unlock()
}

func main() {
	s := SafeCounter{
		NumMap: make(map[string]int),
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("i: ", i)
			s.Add(i)
		}(i)
	}

	wg.Wait()
	fmt.Println(s.NumMap["key"])
}
