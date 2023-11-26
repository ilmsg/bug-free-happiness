package main

import (
	"fmt"
	m "math"
	"os"
)

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
	str := "Learn Go!"
	s2 := `A "raw" string literal
	// can include line breaks.`

	g := 'Σ'

	f := 3.1495
	c := 3 + 4i

	var u uint = 7
	var pi float32 = 22. / 7
	n := byte('\n')

	var a4 [4]int
	a5 := [...]int{3, 1, 5, 10, 100}

	a4_cpy := a4
	a4_cpy[0] = 25
	fmt.Println(a4_cpy[0] == a4[0])

	s3 := []int{4, 5, 9}
	s4 := make([]int, 4)
	var d2 [][]float64
	bs := []byte("A slice")

	s3_cpy := s3
	s3_cpy[0] = 0
	fmt.Println(s3_cpy[0] == s3[0])

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	s = append(s, []int{7, 8, 9}...)
	fmt.Println(s)

	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a5, s4, bs

	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "this is how you write to a file, by the way")
	file.Close()

	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl()
}

func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return // z is implicit here, because we named it earlier.
}

func learnFlowControl() {
	if true {
		fmt.Println("told ya")
	}

	if false {
		// Pout
	} else {
		// Gloat
	}

	x := 42.0
	switch x {
	case 0:
	case 1, 2:
	case 42:
		// Cases don't "fall through."
	case 43:
		// Unreached.
	default:
		// Default case is optional.
	}

	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println(c, "is a string")
	case int64:
		fmt.Printf("%d is an int64\n", c)
	default:
		// all other cases
	}

	for x := 0; x < 3; x++ {
		fmt.Println("iteration", x)
	}

	for {
		break
		continue
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}

	if y := expensiveComputation(); y > x {
		x = y
	}

	xBig := func() bool {
		return x > 100000
	}

	x = 999999
	fmt.Println("xBig:", xBig())

	x = 1.3e3
	fmt.Println("xBig:", xBig())

	fmt.Println("Add + double to number: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2),
	)

	goto love

love:
	learnFunctionFactory()
	learnDefer()
	learnInterfaces()
}

func learnFunctionFactory() {
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautify", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))
}

func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after)
	}
}

func learnDefer() (ok bool) {
	defer fmt.Println("defered statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")
	return true
}

type Stringer interface {
	String() string
}

type pair struct {
	x, y int
}

type Animal interface {
	Say() string
}

type cat struct {
	name string
}

func (c cat) Say() string { return "Meaw" }

func pet() {
	c := cat{"Bob"}
	var i Animal
	i = c
	fmt.Println(i.Say())
}

func (p pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	p := pair{3, 4}
	fmt.Println(p.String())

	var i Stringer
	i = p
	fmt.Println(i.String())

	fmt.Println(p)
	fmt.Println(i)

	learnVariadicParams("great", "learning", "here!")
}

func learnVariadicParams(myString ...interface{}) {
	for _, param := range myString {
		fmt.Println("param:", param)
	}

	fmt.Println("params:", fmt.Sprintln(myString...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// ", ok" idiom used to tell if something worked or not.
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok {
		fmt.Println("on one there")
	} else {
		fmt.Println(x)
	}

	// An error value communicates not just "ok" but
}

func learnMemory() (p, q *int) {
	p = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func expensiveComputation() float64 {
	return m.Exp(10)
}
