package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		fmt.Println("Hi there!")
		<-ticker.C
	}
}

// func main() {
// 	for {
// 		fmt.Println("Hi there!")
// 		time.Sleep(1 * time.Second)
// 	}
// }
