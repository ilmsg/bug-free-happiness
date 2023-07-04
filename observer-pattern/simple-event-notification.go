//go:build ignore

package main

import "sync"

type Subject struct {
	observers []chan string
}

func (s *Subject) Subscribe(observer chan string) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) NotifyObservers(message string) {
	for _, observer := range s.observers {
		observer <- message
	}
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]chan string, 0),
	}
}

func main() {
	subject := NewSubject()

	observer1 := make(chan string)
	observer2 := make(chan string)

	subject.Subscribe(observer1)
	subject.Subscribe(observer2)

	var wg sync.WaitGroup
	wg.Add(2)

}
