package kraken

import (
  "fmt"
)

type addOrderResponse struct {
  Error  []string
}

func (client *Client) Buy(price, amount float64) (err error) {
  fmt.Printf("Buy %.8f at %.2f for %.2f\n", amount, price, amount*price)
  return client.addOrder("buy", price, amount)
}

func (client *Client) Sell(price, amount float64) (err error) {
  fmt.Printf("Sell %.8f at %.2f for %.2f\n", amount, price, amount*price)
  return client.addOrder("sell", price, amount)
}

func (client *Client) addOrder(buySell string, price, amount float64) (err error) {
  params := client.createParams()
  params.Set("pair", client.currencyPair)
  params.Set("type", buySell)
  params.Set("ordertype", "limit")
  params.Set("price", fmt.Sprintf("%.5f", price))
  params.Set("volume", fmt.Sprintf("%.5f", amount))
  if buySell == "sell" {
    params.Set("oflags", "plbc")
  }
  if client.dryRun {
    params.Set("validate", "true")
  }

  var resp addOrderResponse
  err = client.postRequest(API_ADD_ORDER, params, &resp)
  return
}
