package bitfinex

import (
  "bitcoin"
  "errors"
  "strconv"
  "time"
)

type balanceResponse struct {
  Type     string
  Currency string
  Amount   string
}

type balances struct {
  response  []balanceResponse
  timestamp time.Time
}

var balancesCache balances

const (
  BALANCES_CACHE_TIMEOUT = 5 * time.Second
)

func (client *Client) Balance(currency bitcoin.Currency) (
  balance float64, err error) {

  response, err := client.balances()
  if err != nil {
    return
  }
  for _, entry := range response {
    if entry.Type == "exchange" {
      if currency == bitcoin.USD && entry.Currency == "usd" {
        balance, err = strconv.ParseFloat(entry.Amount, 64)
        return
      }
      if currency == bitcoin.BTC && entry.Currency == "btc" {
        balance, err = strconv.ParseFloat(entry.Amount, 64)
        return
      }
    }
  }
  err = errors.New("Currency not found")
  return
}

func (client *Client) balances() (
  balances []balanceResponse, err error) {

  now := time.Now()
  if now.Sub(balancesCache.timestamp) < BALANCES_CACHE_TIMEOUT {
    balances = balancesCache.response
    return
  }

  params := client.createParams()
  err = client.postRequest(API_BALANCE, params, &balances)
  if err != nil {
    return
  }

  balancesCache.response = balances
  balancesCache.timestamp = now
  return
}
