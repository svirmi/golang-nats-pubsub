package db

import (
	"log"

	"github.com/nats-io/nats.go"
	bt "github.com/trever-io/go-blocktrade"
)

type NatsClient struct {
	connection        *nats.Conn
	encodedConnection *nats.EncodedConn
	BidAskTickChan    chan *bt.TickerData
}

func NewNatsClient() (*NatsClient, error) {

	nc, err := nats.Connect("nats://localhost:4222")

	if err != nil {
		log.Fatal(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		log.Fatal(err)
	}

	return &NatsClient{
		connection:        nc,
		encodedConnection: ec,
		BidAskTickChan:    make(chan *bt.TickerData, 1),
	}, nil
}

func (nats *NatsClient) Close() {
	nats.encodedConnection.Close()
	nats.connection.Close()
	log.Println("NATS connection closed")
}
