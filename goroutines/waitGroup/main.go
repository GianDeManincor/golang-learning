package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go callDataBase(&wg)
	go callApi(&wg)
	go processInternal(&wg)

	wg.Wait()
}

func callDataBase(wg *sync.WaitGroup) {
	fmt.Println("callDataBase")
	wg.Done()
}

func callApi(wg *sync.WaitGroup) {
	fmt.Println("callApi")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	fmt.Println("processInternal")
	wg.Done()
}
