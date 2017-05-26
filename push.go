package goPoloniex

import (
	"fmt"

	"gopkg.in/jcelliott/turnpike.v2"
	"strconv"
)

const (
	URL   = "wss://api.poloniex.com"
	REALM = "realm1"
)

type Ticker struct {
	Pair          string
	Last          float64
	LowestAsk     float64
	HighestBid    float64
	PercentChange float64
	BaseVolume    float64
	QuoteVolume   float64
	IsFrozen      float64
	DayHigh       float64
	DayLow        float64
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

	pair := args[0].(string)
	last, _ := strconv.ParseFloat(args[1].(string), 64)
	lowestAsk, _ := strconv.ParseFloat(args[2].(string), 64)
	highestBid, _ := strconv.ParseFloat(args[3].(string), 64)
	percentChange, _ := strconv.ParseFloat(args[4].(string), 64)
	baseVolume, _ := strconv.ParseFloat(args[5].(string), 64)
	quoteVolume, _ := strconv.ParseFloat(args[6].(string), 64)
	isFrozen :=  args[7].(float64)
	dayHigh, _ := strconv.ParseFloat(args[8].(string), 64)
	dayLow, _ := strconv.ParseFloat(args[9].(string), 64)

	t := Ticker{
		pair,
		last,
		lowestAsk,
		highestBid,
		percentChange,
		baseVolume,
		quoteVolume,
		isFrozen,
		dayHigh,
		dayLow,
	}


	return t
}
