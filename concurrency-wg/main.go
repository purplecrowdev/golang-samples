package main

import (
	"fmt"
	"sync"
	"time"
)

func runForSomeTime(rounds int) {
	for i := 0; i < rounds; i++ {
		time.Sleep(time.Second)
		fmt.Println("iteration", i)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		runForSomeTime(10)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		runForSomeTime(20)
		wg.Done()
	}()

	wg.Wait()
}
