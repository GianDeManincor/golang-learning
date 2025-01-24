package model

type Automobile struct {
	Age   int
	Plate string
	Brand string
}

type Motocycle struct {
	Automobile
	CylinderCapacity int
}

type Car struct {
	Automobile
	NumbersOfDoors int
	Potency        int
}
