package main

import (
	"errors"
	"log"
)

type Queue interface {
	IsEmpty() bool
	GetLength() int
	Dequeue() interface{}
	Enqueue(element interface{})
	GetElements() []interface{}
}

type queue struct {
	Elements []interface{}
	Size     int
}

func (q *queue) Enqueue(element interface{}) {
	if q.GetLength() == q.Size {
		log.Fatal("Overflow")
		return
	}
	q.Elements = append(q.Elements, element)
}

func (q *queue) Dequeue() interface{} {
	if q.IsEmpty() {
		log.Println("underflow")
		return 0
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element
}

func (q *queue) GetLength() int {
	return len(q.Elements)
}

func (q *queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

func (q *queue) GetElements() []interface{} {
	return q.Elements
}

func NewQueue(size int) Queue {
	return &queue{
		Elements: make([]interface{}, 0),
		Size:     size,
	}
}
