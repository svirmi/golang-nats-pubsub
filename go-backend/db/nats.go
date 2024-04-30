package db

import (
	"log"

	"github.com/nats-io/nats.go"
	bt "github.com/trever-io/go-blocktrade"
)

type NatsClient struct {
	encodedConnection *nats.EncodedConn
	BidAskTickChan    chan *bt.TickerData
}

func NewClient() (*NatsClient, error) {

	nc, err := nats.Connect("nats://nats:4222")

	if err != nil {
		log.Fatal(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		nc.Close()
		nc.Close()
		log.Println("NATS connection closed")
	}()

	return &NatsClient{
		encodedConnection: ec,
		BidAskTickChan:    make(chan *bt.TickerData, 1),
	}, nil
}
