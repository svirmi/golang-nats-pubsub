package main

import (
	"fmt"
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

	stream := make(chan Data)

	go func() {
		interrupt := make(chan os.Signal, 1)
		<-interrupt

		fmt.Println("finishing program")

		close(stream)
	}()

	fmt.Println("hello world!")
}
