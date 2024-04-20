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

	randStream := repeatFunc(interrupt, randNumFetcher)

	for rando := range take(interrupt, randStream, 10) {
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

func take[T any](done <-chan os.Signal, stream <-chan T, n int) <-chan T {

	// Send operation to unbuffered channel blocks the sending goroutine,
	// 'stream' (see repeatFunc, it stops) in this case blocked until data is read by 'taken' channel

	taken := make(chan T)

	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}
