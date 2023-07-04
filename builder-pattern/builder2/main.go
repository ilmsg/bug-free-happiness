package main

import "fmt"

func main() {
	electricCarBuilder := &ElectricCarBuilder{}
	gasolineCarBuilder := &GasolineCarBuilder{}

	electricCar := CreateCar(electricCarBuilder, "Carolla")
	gasolineCar := CreateCar(gasolineCarBuilder, "Carolla")

	fmt.Printf("Electric car: %+v\n", electricCar)
	fmt.Printf("Gasoline car: %+v\n", gasolineCar)
}
