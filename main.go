package main

import (
	"fmt"

	"sync"

	"github.com/branda22/hello/poloniex"
)

func ondataReceived(args []interface{}, kwargs map[string]interface{}) {
	fmt.Println(args)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Poloniex Market Data Stream V1.0")

	wg.Add(1)

	poloniexClient := poloniex.NewPoloniexClient()

	poloniexClient.SubscribeTicker(ondataReceived)
	wg.Wait()
}
