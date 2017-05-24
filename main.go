package main

import (
	"fmt"

	"sync"

	"github.com/branda22/goPoloniex/poloniex"
)

func ondataReceived(args []interface{}, kwargs map[string]interface{}) {
	fmt.Println(args)
}

func onTrollMessage(args []interface{}, kwargs map[string]interface{}) {
	fmt.Println(args)
}

func retrieveMessage(dataChannel chan []interface{}) {
	fmt.Println("retrieveMessage Invoked")
	for {
		message := <-dataChannel

		if len(message) > 0 {
			fmt.Println(message[0], message[1])
		}
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Poloniex Market Data Stream V1.0")

	wg.Add(1)

	pushClient := poloniex.NewPushClient()

	//pushClient.SubscribeTicker(ondataReceived)
	//pushClient.SubscribeOrderBookTrades("USDT_BTC", onTrollMessage)

	dataChannel := make(chan []interface{})

	err := pushClient.ChannelTicker(dataChannel)

	if err != nil {
		panic(err)
	}

	go retrieveMessage(dataChannel)
	wg.Wait()
}
