package main

import "fmt"

func main() {
	// criando channel para receber mensagem
	messages := make(chan string)

	// colocando mensagem no channel
	go func() {
		messages <- "Olá mundo!"
		close(messages) // fechando o channel após colocar a mensagem para evitar deadlock
	}()

	// lendo a mensagem
	greeting := <-messages
	fmt.Println(greeting)
}
