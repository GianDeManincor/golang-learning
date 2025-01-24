package model

import "time"

type Person struct {
	Name        string
	Address     Address
	DateOfBirth time.Time
	Age         int
}

// func (p *Person) CalculateAge() {
// 	ageOfBirth := p.DateOfBirth.Year()
// 	currentAge := time.Now().Year()
// 	p.Age = currentAge - ageOfBirth
// }

func (p Person) CalculateAge() int {
	ageOfBirth := p.DateOfBirth.Year()
	currentAge := time.Now().Year()
	return currentAge - ageOfBirth
}

func CalculateAge(p *Person) int {
	ageOfBirth := p.DateOfBirth.Year()
	currentAge := time.Now().Year()
	return currentAge - ageOfBirth
}
