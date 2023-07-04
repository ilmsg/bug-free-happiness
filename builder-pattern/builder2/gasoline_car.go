package main

type GasolineCarBuilder struct {
	Car
}

func (b *GasolineCarBuilder) setModel(model string) CarBuilder {
	b.Model = model
	return b
}

func (b *GasolineCarBuilder) Build() Car {
	return b.Car
}
