package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

type Data struct {
	ID    int
	Value string
}

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	defer func() {
		close(interrupt)
	}()

	randNumFetcher := func() int {
		return rand.Intn(500000000)
	}

	for rando := range repeatFunc(interrupt, randNumFetcher) {
		fmt.Println(rando)
	}

	fmt.Println("\nProgram finished")
}

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}
