package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado")
		}
	}()

	_, err := os.Open("./settings.txt")
	if err != nil {
		panic("File not exist")
	}
}

func main() {
	ReadFile()
	fmt.Println("Fim")
}
