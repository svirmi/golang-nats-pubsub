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
	primeStream := primeFinder(interrupt, randStream)

	for rando := range take(interrupt, primeStream, 10) {
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

func primeFinder(done <-chan os.Signal, randStream <-chan int) <-chan int {
	// NB! very costly slow function
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)

	go func() {

		defer close(primes)

		for {
			select {
			case <-done:
				return
			case randomInt := <-randStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}
