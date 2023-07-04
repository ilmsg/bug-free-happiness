package main

type Car struct {
	Model string
}

type CarBuilder interface {
	setModel(model string) CarBuilder
	Build() Car
}

func CreateCar(builder CarBuilder, modelName string) Car {
	return builder.setModel(modelName).Build()
}
