package btce

import (
  "bitcoin"
  "errors"
  "time"
)

type Info struct {
  Success   int
  Return    InfoReturn
  timestamp time.Time
}

type InfoReturn struct {
  Funds InfoFunds
}

type InfoFunds struct {
  Usd float64
  Btc float64
}

const (
  GET_INFO_CACHE_TIMEOUT = 5 * time.Second
)

func (client *Client) GetInfo() (info Info, err error) {
  now := time.Now()
  if now.Sub(client.infoCache.timestamp) < GET_INFO_CACHE_TIMEOUT {
    info = client.infoCache
    return
  }

  params := client.createParams()
  err = client.postRequest(API_GET_INFO, params, &info)
  info.timestamp = now
  client.infoCache = info
  return
}

func (client *Client) Balance(
  currency bitcoin.Currency) (balace float64, err error) {

  info, err := client.GetInfo()
  if err != nil {
    return
  }

  if currency == bitcoin.USD {
    balace = info.Return.Funds.Usd
  } else if currency == bitcoin.BTC {
    balace = info.Return.Funds.Btc
  } else {
    err = errors.New("Unsupported currency")
  }
  return
}

func (client *Client) Fee() (fee float64, err error) {
  return 0.002, nil
}
