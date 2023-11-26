package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func main() {
	start := time.Now()
	resch := make(chan string)

	wg = &sync.WaitGroup{}
	wg.Add(3)

	go doWork(2*time.Second, resch)
	go doWork(4*time.Second, resch)
	go doWork(6*time.Second, resch)

	go func() {
		for res := range resch {
			fmt.Println(res)
		}
	}()

	wg.Wait()
	close(resch)
	fmt.Printf("work took %v seconds\n", time.Since(start))
	// time.Sleep(time.Second)
}

func doWork(d time.Duration, ch chan string) {
	defer wg.Done()
	time.Sleep(d)
	ch <- fmt.Sprintf("the result of the work -> %d", rand.Intn(100))
}
