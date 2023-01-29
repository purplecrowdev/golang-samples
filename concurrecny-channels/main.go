package main

import (
	"fmt"
	"time"
)

func runForSomeTime(rounds int, ch chan int) {

	for i := 0; i < rounds; i++ {
		time.Sleep(time.Second)
		ch <- i
	}

	close(ch)
}

func main() {

	ch := make(chan int)

	go runForSomeTime(10, ch)

	for x := range ch {
		fmt.Println(x)
	}

}
