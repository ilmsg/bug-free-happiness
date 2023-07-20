package main

import "fmt"

func main() {
	storage := make(map[string]int)

	storage["key1"] = 1001
	storage["key2"] = 1002

	fmt.Printf("key1: %d\n", storage["key1"])
	fmt.Printf("key2: %v\n", storage["key2"])

}
