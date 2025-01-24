package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	largura, altura float64
}

type circulo struct {
	radius float64
}

func ExibirAreaRetangulo(retangulo retangulo) {
	area := retangulo.largura * retangulo.altura
	fmt.Println(area)
}

func ExibirAreaCirculo(circulo circulo) {
	area := math.Pi * circulo.radius * circulo.radius
	fmt.Println(area)
}

func main() {
	fmt.Println("Initing...")

	retangulo := retangulo{
		largura: 1,
		altura:  2,
	}

	circulo := circulo{
		radius: 3,
	}

	ExibirAreaRetangulo(retangulo)
	ExibirAreaCirculo(circulo)
}
