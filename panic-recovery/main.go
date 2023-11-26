package main

import "fmt"

func main() {
	fmt.Println("start")
	performDevision(10, 2)
	performDevision(8, 0)
	performDevision(12, 3)
	fmt.Println("end")
}

func performDevision(a, b int) {
	defer recoveryFromPanic()
	// if b == 0 {
	// 	panic("Division by zero")
	// }
	result := a / b
	fmt.Println("Result:", result)
}

func recoveryFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
