package bitstamp

import (
  "bitcoin"
  "errors"
  "time"
)

type Balance struct {
  Usd          float64
  Btc          float64
  UsdReserved  float64
  BtcReserved  float64
  UsdAvailable float64
  BtcAvailable float64
  Fee          float64
  timestamp    time.Time
}

const (
  BALANCE_CACHE_TIMEOUT = 5 * time.Second
)

func (client *Client) BitstampBalance() (balance Balance, err error) {
  now := time.Now()
  if now.Sub(client.balanceCache.timestamp) < BALANCE_CACHE_TIMEOUT {
    balance = client.balanceCache
    return
  }

  params := client.createParams()
  result, err := requestMap(API_BALANCE, params)
  if err != nil {
    return
  }
  balance = Balance{
    Usd:          result.getFloat("usd_balance"),
    Btc:          result.getFloat("btc_balance"),
    UsdReserved:  result.getFloat("usd_reserved"),
    BtcReserved:  result.getFloat("btc_reserved"),
    UsdAvailable: result.getFloat("usd_available"),
    BtcAvailable: result.getFloat("btc_available"),
    Fee:          result.getFloat("fee"),
    timestamp:    now,
  }
  client.balanceCache = balance
  return
}

func (client *Client) Balance(currency bitcoin.Currency) (
  balance float64, err error) {

  if currency != bitcoin.USD && currency != bitcoin.BTC {
    err = errors.New("Unsupported currency")
    return
  }
  balances, err := client.BitstampBalance()
  if err != nil {
    return
  }
  if currency == bitcoin.USD {
    balance = balances.Usd
  }
  if currency == bitcoin.BTC {
    balance = balances.Btc
  }
  return
}

func (client *Client) Fee() (fee float64, err error) {
  balances, err := client.BitstampBalance()
  if err != nil {
    return
  }
  fee = balances.Fee / 100
  return
}
