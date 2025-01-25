package main

import (
	"fmt"
	"os"
)

func exampleFuncDefer() {
	fmt.Println("exemplo func defer")
}

func main() {
	file, err := os.Create("./settings.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer exampleFuncDefer()

	fmt.Println("Gravando texto no arquivo:", file.Name())

	_, err = file.Write([]byte("teste"))
	if err != nil {
		panic(err)
	}
}
