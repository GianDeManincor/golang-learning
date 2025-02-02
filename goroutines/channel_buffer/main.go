package main

import (
	"fmt"
	"time"
)

func main() {
	// criando channel para receber mensagem com quantidade de buffer
	numbers := make(chan int, 10)

	// colocando mensagem no channel
	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
			fmt.Printf("Enviando %d! \n", i)
		}
		close(numbers) // fechando o channel após colocar a mensagem para evitar deadlock
	}()

	// lendo a mensagem com buffer
	for v := range numbers {
		fmt.Printf("Recebendo %d! \n", v)
		time.Sleep(time.Second)
	}

}
