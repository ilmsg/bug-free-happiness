package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered 1"
	messages <- "channel 1"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages <- "buffered 2"
	messages <- "channel 2"

	close(messages)
	for msg := range messages {
		fmt.Println(msg)
	}
}
