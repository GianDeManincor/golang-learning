package main

import (
	"fmt"
	"time"

	"github.com/giandemanincor/encapsulation/model"
)

func main() {
	address := model.Address{
		Street: "Rua Duarte Pacheco",
		Number: 1400,
		City:   "São José do Rio Preto",
	}

	person := model.Person{
		Name:        "Gian de Manincor",
		Address:     address,
		DateOfBirth: time.Date(1995, 07, 26, 0, 0, 0, 0, time.Local),
	}

	automobileCar := model.Automobile{
		Age:   2024,
		Plate: "SWA6B03",
		Brand: "Tiggo 5X - Caoa Cherry",
	}

	car := model.Car{
		Automobile:     automobileCar,
		NumbersOfDoors: 4,
		Potency:        150,
	}

	fmt.Println(person)

	// Chamando o método CalculateAge
	age := person.CalculateAge()
	fmt.Println("Age with method:", age)

	// Chamando a função CalculateAge
	age = model.CalculateAge(&person)
	fmt.Println("Age with func:", age)

	fmt.Println(car)

	fmt.Println(model.Initializer())
}
