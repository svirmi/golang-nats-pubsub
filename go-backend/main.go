package main

import (
	"fmt"
	"log"

	bt "github.com/trever-io/go-blocktrade"
)

func main() {

	blocktradeClient := bt.NewClient("", "")

	assets, err := blocktradeClient.TradingAssets()

	if err != nil {
		log.Fatal(err)
	}

	for i, asset := range assets {
		fmt.Println(i, asset.Id, asset.FullName)
	}
}
