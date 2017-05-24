package poloniex

type PublicClient struct{}

func NewPublicClient() PublicClient {
	return PublicClient{}
}
//
//func (c *PublicClient) ReturnTicker(currencyPair string) []interface{} {
//
//}
//
//func (c *PublicClient) ReturnOrderBook(currencyPair string) map[string]interface{} {
//
//}
//
//func (c *PublicClient) ReturnTradeHistory(currencyPair string) []map[string]interface{} {
//
//}
//
//func (c *PublicClient) ReturnChartData(currencyPair string) []map[string]interface{} {
//
//}
//
//func (c *PublicClient) ReturnCurrencies() map[string]interface{} {
//
//}
//
//func (c *PublicClient) ReturnLoanOrders(currencyPair string) map[string]interface{} {
//
//}
