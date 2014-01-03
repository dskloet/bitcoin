package btce

import (
  "errors"
  "fmt"
)

type TradeResponse struct {
  Success int
  Error   string
}

func (client *Client) trade(tradeType string, price, amount float64) (err error) {
  params := client.createParams()
  params["pair"] = []string{"btc_usd"}
  params["type"] = []string{tradeType}
  params["rate"] = []string{fmt.Sprintf("%.3f", price)}
  params["amount"] = []string{fmt.Sprintf("%.8f", amount)}
  var resp TradeResponse
  err = client.postRequest(API_TRADE, params, &resp)
  if err != nil {
    return
  }
  if resp.Success == 0 {
    err = errors.New(resp.Error)
  }
  return
}

func (client *Client) Buy(price, amount float64) (err error) {
  return client.trade("buy", price, amount)
}

func (client *Client) Sell(price, amount float64) (err error) {
  return client.trade("sell", price, amount)
}
