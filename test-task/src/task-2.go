package main

import (
	"fmt"
	"math/rand"
)

func performTask2() {
	results := make(chan string) // here I make the execution slower by not making the channel of size 20
	cancel := make(chan struct{})
	aGoroutine := func() {
		for i := 0; i < 10; i++ {
			select {
			case results <- fmt.Sprintf("%d", i):
			case <-cancel:
				return
			}
		}
	}
	go aGoroutine()
	go aGoroutine()
	for i := 0; i < 20; i++ {
		<-results
		if rand.Intn(42) < 24 {
			close(cancel)
			break
		}
	}
}
