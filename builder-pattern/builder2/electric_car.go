package main

type ElectricCarBuilder struct {
	Car
}

func (b *ElectricCarBuilder) setModel(model string) CarBuilder {
	b.Model = model
	return b
}

func (b *ElectricCarBuilder) Build() Car {
	return b.Car
}
