package poloniex

import (
	"fmt"

	"gopkg.in/jcelliott/turnpike.v2"
)

const (
	URL   = "wss://api.poloniex.com"
	REALM = "realm1"
)

type PoloClient struct {
	WSClient *Client
}

func NewPoloniexClient() PoloClient {
	wampClient, err := turnpike.NewWebsocketClient(turnpike.JSON, URL, nil, nil)

	if err != nil {
		fmt.Errorf("Error creating turnpike websocket client")
	}

	_, err = c.WSClient.JoinRealm(REALM, nil)

	if err != nil {
		fmt.Errorf("Error joining realm")
	}

	return Client{
		WSClient: wampClient,
	}
}

func (c *PoloClient) SubscribeTicker(callback func(args []interface{}, kwargs map[string]interface{})) error {

	err = c.WSClient.Subscribe("ticker", nil, callback)

	if err != nil {
		return err
	}

	return nil
}

func (c *PoloClient) SubscribeOrderBookTrades(currencyPair string, callback func(args []interface{}, kwargs map[string]interface{})) error {

	err = c.WSClient.Subscribe(currencyPair, nil, callback)

	if err != nil {
		return err
	}

	return nil
}

func (c *PoloClient) SubscribeTrollbox(callback func(args []interface{}, kwargs map[string]interface{})) error {

	err = c.WSClient.Subscribe("trollbox", nil, callback)

	if err != nil {
		return err
	}

	return nil
}
