package bitfinex

import (
  "errors"
  "fmt"
)

func (client *Client) newOrder(side string, price, amount float64) (err error) {
  params := client.createParams()
  params["symbol"] = client.currencyPair
  params["side"] = side
  params["price"] = fmt.Sprintf("%.5f", price)
  params["amount"] = fmt.Sprintf("%.8f", amount)
  params["exchange"] = "bitfinex"
  params["type"] = "exchange limit"
  var resp map[string]interface{}
  err = client.postRequest(API_ORDER_NEW, params, &resp)
  if err != nil {
    return
  }
  if message, ok := resp["message"]; ok {
    err = errors.New(message.(string))
  }
  return
}

func (client *Client) Buy(price, amount float64) (err error) {
  return client.newOrder("buy", price, amount)
}

func (client *Client) Sell(price, amount float64) (err error) {
  return client.newOrder("sell", price, amount)
}
