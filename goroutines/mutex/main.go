package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var m sync.Mutex
	count := 0

	for i := 0; i < 10000; i++ {
		go func() {
			m.Lock()
			count++
			m.Unlock()
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println(count)
}
