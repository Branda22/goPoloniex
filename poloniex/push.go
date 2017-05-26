package poloniex

import (
	"fmt"

	"gopkg.in/jcelliott/turnpike.v2"
)

const (
	URL   = "wss://api.poloniex.com"
	REALM = "realm1"
)

type Ticker struct {
	Pair          string
	Last          float32
	LowestAsk     float32
	HighestBid    float32
	PercentChange float32
	BaseVolume    float32
	QuoteVolume   float32
	IsFrozen      int
	DayHigh       float32
	DayLow        float32
}

type PushClient struct {
	WSClient *turnpike.Client
}

func NewPushClient() PushClient {
	wampClient, err := turnpike.NewWebsocketClient(turnpike.JSON, URL, nil, nil)

	if err != nil {
		fmt.Errorf("Error creating turnpike websocket client")
	}

	_, err = wampClient.JoinRealm(REALM, nil)

	if err != nil {
		fmt.Errorf("Error joining realm")
	}

	return PushClient{
		wampClient,
	}
}

func (c *PushClient) SubscribeTicker(callback func(args []interface{}, kwargs map[string]interface{})) error {

	err := c.WSClient.Subscribe("ticker", nil, callback)

	if err != nil {
		return err
	}

	return nil
}

func (c *PushClient) SubscribeOrderBookTrades(currencyPair string, callback func(args []interface{}, kwargs map[string]interface{})) error {

	err := c.WSClient.Subscribe(currencyPair, nil, callback)

	if err != nil {
		return err
	}

	return nil
}

func (c *PushClient) SubscribeTrollbox(callback func(args []interface{}, kwargs map[string]interface{})) error {

	err := c.WSClient.Subscribe("trollbox", nil, callback)

	if err != nil {
		return err
	}

	return nil
}

/* Channel API */
func (c *PushClient) ChannelTicker(dataChannel chan Ticker) error {
	err := c.WSClient.Subscribe("ticker", nil, func(args []interface{}, kwargs map[string]interface{}) {
		dataChannel <- parseTicker(args)
	})

	if err != nil {
		return err
	}

	return nil
}

func parseTicker(args []interface{}) Ticker {
	t := Ticker{
		args[0].(string),
		args[1].(float32),
		args[2].(float32),
		args[3].(float32),
		args[4].(float32),
		args[5].(float32),
		args[6].(float32),
		args[7].(int),
		args[8].(float32),
		args[9].(float32),
	}

	return t
}
