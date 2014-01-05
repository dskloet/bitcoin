package kraken

import (
  "bitcoin"
  "fmt"
)

type balanceResponse struct {
  Error  []string
  Result map[string]float64
}

func (client *Client) Balance(
  currency bitcoin.Currency) (balance float64, err error) {
  if currency == bitcoin.BTC {
    return
  }

  params := client.createParams()
  var resp balanceResponse
  err = client.postRequest(API_BALANCE, params, &resp)
  if err != nil {
    return
  }
  fmt.Printf("resp = %v\n", resp)
  balance = resp.Result["USD"]
  return
}
