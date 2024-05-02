package main

import (
	"backend/db"
	"fmt"
	"log"

	bt "backend/blocktrade"
)

func main() {

	natsClient, err := db.NewNatsClient()

	if err != nil {
		log.Fatal(err)
	}

	blocktradeClient := bt.NewClient("", "")

	defer func() {
		natsClient.Close()
		blocktradeClient.Close()

		log.Println("WS connection closed")
	}()

	assets, err := blocktradeClient.TradingAssets()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trading assets:")

	for i, asset := range assets {
		fmt.Println(i, asset.Id, asset.FullName)
	}

	pairs, err := blocktradeClient.TradingPairs()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trading pairs:")

	for i, pair := range pairs {
		fmt.Println(i, pair.Id, pair.BaseAssetId, pair.QuoteAssetId)
	}

	_, err = blocktradeClient.Websocket()

	if err != nil {
		log.Fatal(err)
	}

	stream := func(TickerResponse *bt.TickerResponse, err error) {
		fmt.Println(TickerResponse.Data)
	}

	err = blocktradeClient.SubscribeTicker(14, stream)

	if err != nil {
		log.Fatal(err)
	}

	select {}

}
