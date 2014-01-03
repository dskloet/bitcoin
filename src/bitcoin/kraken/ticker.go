package kraken

import (
  "strconv"
)

type tickerResponse struct {
  Error  []string
  Result map[string]tickerResult
}

type tickerResult struct {
  C []string
}

func (client Client) LastPrice() (price float64, err error) {
  var resp tickerResponse
  err = client.getRequest(API_TICKER + client.currencyPair, &resp)
  if err != nil {
    return
  }
  return strconv.ParseFloat(resp.Result[client.currencyPair].C[0], 64)
}
