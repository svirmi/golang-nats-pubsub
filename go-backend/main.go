package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Data struct {
	ID    int
	Value string
}

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	stream := make(chan Data, 1)  // Channel to communicate between goroutines
	closer := make(chan struct{}) // Channel to signal to stop generation and exit goroutine

	go func() {
		<-interrupt
		close(closer)
	}()

	var wg sync.WaitGroup // WaitGroup to synchronize goroutines

	// Start concurrent generation of data
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			generateData(id, stream, closer)
		}(i)
	}

	// Print data received from the stream
	go func() {
		for data := range stream {
			fmt.Printf("Received: %+v\n", data)
		}
	}()

	wg.Wait()     // Wait for all goroutines to finish
	close(stream) // Close the stream channel

	fmt.Println("Program finished")
}

func generateData(id int, stream chan<- Data, closer chan struct{}) {
loop:
	for i := 0; i < 3; i++ {
		data := Data{
			ID:    id,
			Value: fmt.Sprintf("Data %d-%d", id, i),
		}

		select {
		case stream <- data: // Send data through the stream channel
		case <-closer: // Exit loop and stop sending/generation
			break loop
		}
	}
}
