package bitstamp

import (
  "bitcoin"
  "errors"
  "time"
)

type Ticker struct {
  Last      float64
  High      float64
  Low       float64
  Volume    float64
  Bid       float64
  Ask       float64
  timestamp time.Time
}

const (
  TICKER_CACHE_TIMEOUT = 5 * time.Second
)

func (client Client) Ticker() (ticker Ticker, err error) {
  now := time.Now()
  if now.Sub(client.tickerCache.timestamp) < TICKER_CACHE_TIMEOUT {
    ticker = client.tickerCache
    return
  }

  result, err := getMap(API_TICKER)
  if err != nil {
    return
  }
  return Ticker{
    Last:      result.getFloat("last"),
    High:      result.getFloat("high"),
    Low:       result.getFloat("low"),
    Volume:    result.getFloat("volume"),
    Bid:       result.getFloat("bid"),
    Ask:       result.getFloat("ask"),
    timestamp: now,
  }, nil
}

func (client Client) LastRate(first, second bitcoin.Currency) (
  rate float64, err error) {

  if first == second {
    rate = 1
    return
  }
  ticker, err := client.Ticker()
  if err != nil {
    return
  }
  if first == bitcoin.BTC && second == bitcoin.USD {
    rate = ticker.Last
  } else if first == bitcoin.USD && second == bitcoin.BTC {
    rate = 1 / ticker.Last
  } else {
    err = errors.New("Currency combination not supported")
  }
  return
}
