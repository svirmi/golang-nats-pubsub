package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"

	"github.com/nats-io/nats.go"
)

type SensorData struct {
	SensorID int
	Value    int
}

var (
	nc   *nats.Conn
	subj = "sensorData"
	err  error
)

func main() {
	numSensors := 1000
	numWorkers := runtime.NumCPU() * 4

	dataStream := make(chan SensorData) // Channel for data streams
	done := make(chan struct{})         // Channel to signal when processing is done

	nc, err = nats.Connect("nats://nats:4222")

	if err != nil {
		log.Fatal(err)
	}

	// Create worker pool
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataStream, &wg, done)
	}

	// Generate data streams
	for i := 0; i < numSensors; i++ {
		go func(sensorID int) {
			for {
				data := SensorData{
					SensorID: sensorID,
					Value:    rand.Intn(5000000), // Random value for demonstration
				}
				dataStream <- data
				// time.Sleep(500 * time.Millisecond) // Data sent every half a second
			}
		}(i)
	}

	// Wait for all data to be processed
	go func() {
		wg.Wait()
		close(dataStream)
		close(done)
	}()

	// Wait for the processing to complete
	<-done
	fmt.Println("All data processed. Exiting.")
}

func worker(id int, dataStream <-chan SensorData, wg *sync.WaitGroup, done chan<- struct{}) {
	defer wg.Done()

	for data := range dataStream {
		// Simulate data processing/publishing
		publishData(id, data)
	}

	fmt.Printf("Worker %d finished\n", id)
	done <- struct{}{}
}

func publishData(workerID int, data SensorData) {

	// NB! very costly slow function
	// finding prime numbers
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	// now := time.Now()

	if isPrime(data.Value) {
		// duration := time.Since(now)
		_ = workerID

		msgBody, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling to bytes : ", err)
		}

		if err := nc.Publish(subj, msgBody); err != nil {
			log.Fatal(err)
		}
		nc.Flush()

		// fmt.Printf("Worker %d processed data from Sensor %d: Prime number is %d, took %d milliseconds to find\n", workerID, data.SensorID, data.Value, duration.Milliseconds())
	}
}
