package main

import (
	"fmt"
	"sync"
)

// Stock is our Subject
type Stock struct {
	price   float64
	brokers []chan float64
}

// NewStock creates a new Stock
func NewStock() *Stock {
	return &Stock{
		// price:   0.0,
		brokers: make([]chan float64, 0),
	}
}

// Register method for adding new broker
func (s *Stock) Register(broker chan float64) {
	s.brokers = append(s.brokers, broker)
}

// UpdatePrice updates the stock price and notifies all brokers
func (s *Stock) UpdatePrice(price float64) {
	s.price = price
	for _, broker := range s.brokers {
		broker <- s.price
	}
}

func main() {
	stock := NewStock()

	// creating channels for brokers
	broker1 := make(chan float64)
	broker2 := make(chan float64)

	stock.Register(broker1)
	stock.Register(broker2)

	var wg sync.WaitGroup // using WaitGroup
	go func() {
		for price := range broker1 {
			fmt.Printf("Broker 1 updated the stock price: %.2f\n", price)
			wg.Done()
		}
	}()

	go func() {
		for price := range broker2 {
			fmt.Printf("Broker 2 updated the stock price: %.2f\n", price)
			wg.Done()
		}
	}()

	wg.Add(2) // assuming each observer will get exactly one message
	stock.UpdatePrice(150.25)
	stock.UpdatePrice(152.85)

	wg.Wait() // wait until all observers are done
}
