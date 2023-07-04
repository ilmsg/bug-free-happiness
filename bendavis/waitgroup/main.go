package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("IDX from first func:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("IDX from second func:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Done.")
}
