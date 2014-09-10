package kraken

import (
  "bitcoin"
  "strconv"
  "time"
)

type balanceResponse struct {
  Error  []string
  Result map[string]string
}

var balanceCache balanceResponse
var balanceCacheTime time.Time

const (
  BALANCES_CACHE_TIMEOUT = 5 * time.Second
)

func (client *Client) Balance(currency bitcoin.Currency) (
    balance float64, err error) {
  resp, err := client.getBalanceResponse()
  if err != nil {
    return
  }

  if currency == bitcoin.BTC {
    balance, _ = strconv.ParseFloat(resp.Result["XXBT"], 64)
  } else if currency == bitcoin.FIAT {
    balance, _ = strconv.ParseFloat(resp.Result["ZEUR"], 64)
  }
  return
}

func (client *Client) getBalanceResponse() (resp balanceResponse, err error) {
  now := time.Now()
  if now.Sub(balanceCacheTime) < BALANCES_CACHE_TIMEOUT {
    resp = balanceCache
    return
  }

  params := client.createParams()
  err = client.postRequest(API_BALANCE, params, &resp)
  if err != nil {
    return
  }

  balanceCache = resp
  balanceCacheTime = now
  return
}
