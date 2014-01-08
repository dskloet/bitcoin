package bitfinex

import (
  "strconv"
)

type tickerResponse struct {
  Last_price string
}

func (client Client) LastPrice() (price float64, err error) {
  var resp tickerResponse
  err = client.getRequest(API_TICKER + client.currencyPair, &resp)
  if err != nil {
    return
  }
  return strconv.ParseFloat(resp.Last_price, 64)
}
