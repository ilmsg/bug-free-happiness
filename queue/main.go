package main

import "fmt"

func main() {
	q := NewQueue(4)

	fmt.Println(q.GetElements())

	q.Enqueue(1)
	fmt.Println(q.GetElements())

	q.Enqueue(2)
	fmt.Println(q.GetElements())

	q.Dequeue()
	fmt.Println(q.GetElements())

	q.Enqueue(4)
	fmt.Println(q.GetElements())

	q.Dequeue()
	fmt.Println(q.GetElements())

	q.Dequeue()
	fmt.Println(q.GetElements())
}
