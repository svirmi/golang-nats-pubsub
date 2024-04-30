package models

import (
	"log/slog"
	"sync"

	bt "github.com/trever-io/go-blocktrade"
)

type Stream struct {
	BidAskTickChan chan *bt.TickerData
	Wg             *sync.WaitGroup
	Closer         chan interface{}
	Logger         *slog.Logger
}
