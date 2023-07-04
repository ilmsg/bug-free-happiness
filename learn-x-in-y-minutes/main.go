package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	beyondHello()
}

func beyondHello() {
	var x int
	x = 3

	y := 4

	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	// str := "Learn Go!"
	// s2 := `A "raw" string literal
	// can include line breaks.`

	// g := 'Σ'

	// f := 3.1495
	// c := 3 + 4i

	// var u uint = 7
	// var pi float32 = 22. / 7
	// n := byte('\n')

	var a4 [4]int
	// a5 := [...]int{3, 1, 5, 10, 100}

	a4_cpy := a4
	a4_cpy[0] = 25
	fmt.Println(a4_cpy[0] == a4[0])

	s3 := []int{4, 5, 9}
	// s4 := make([]int, 4)
	// var d2 [][]float64
	// bs := []byte("A slice")

	s3_cpy := s3
	s3_cpy[0] = 0
	fmt.Println(s3_cpy[0] == s3[0])
}
