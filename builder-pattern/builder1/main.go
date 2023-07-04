package main

import "fmt"

type Car struct {
	Model string
}

type CarBuilder struct {
	Car
}

func (cb *CarBuilder) setModel(model string) *CarBuilder {
	cb.Model = model
	return cb
}

func (cb *CarBuilder) Build() *Car {
	return &cb.Car
}

func main() {
	carBuilder := &CarBuilder{}
	car := carBuilder.
		setModel("Carolla").
		Build()

	fmt.Printf("Model: %s\n", car.Model)
}
