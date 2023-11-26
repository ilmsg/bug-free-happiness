package main

func main() {
	a := 10
	b := 20
	sum := add(a, b)
	println("Sum:", sum)
}

func add(x, y int) int {
	return x + y
}
