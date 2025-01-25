package main

import "os"

func main() {
	_, err := os.Open("f://settings.txt")

	if err != nil {
		panic(err)
	}
}
