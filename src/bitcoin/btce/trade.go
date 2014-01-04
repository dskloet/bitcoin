package btce

import (
  "errors"
  "fmt"
)

type tradeResponse struct {
  Success int
  Error   string
}

func (client *Client) trade(tradeType string, price, amount float64) (err error) {
  if client.dryRun {
    fmt.Printf("Skipped\n")
    return
  }
  params := client.createParams()
  params.Set("pair", "btc_usd")
  params.Set("type", tradeType)
  params.Set("rate", fmt.Sprintf("%.3f", price))
  params.Set("amount", fmt.Sprintf("%.8f", amount))
  var resp tradeResponse
  err = client.postRequest(API_TRADE, params, &resp)
  if err != nil {
    return
  }
  if resp.Success != 1 {
    err = errors.New(resp.Error)
  }
  return
}

func (client *Client) Buy(price, amount float64) (err error) {
  fmt.Printf("Buy %.8f at %.3f for %.3f\n", amount, price, amount*price)
  return client.trade("buy", price, amount)
}

func (client *Client) Sell(price, amount float64) (err error) {
  fmt.Printf("Sell %.8f at %.3f for %.3f\n", amount, price, amount*price)
  return client.trade("sell", price, amount)
}
