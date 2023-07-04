package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < 5; i++ {
			fmt.Println("IDX from first func:", i)
			sum += i
		}
		ch <- sum
	}()

	output := <-ch
	fmt.Println("Output:", output)
}
