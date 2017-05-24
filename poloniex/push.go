package poloniex

import (
	"fmt"

	"gopkg.in/jcelliott/turnpike.v2"
)

const (
	URL   = "wss://api.poloniex.com"
	REALM = "realm1"
)

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
func (c *PushClient) ChannelTicker(dataChannel chan []interface{}) error {
	err := c.WSClient.Subscribe("ticker", nil, func(args []interface{}, kwargs map[string]interface{}) {
		dataChannel <- args
	})

	if err != nil {
		return err
	}

	return nil
}
