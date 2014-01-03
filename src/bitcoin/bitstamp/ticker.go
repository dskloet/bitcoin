package bitstamp

import (
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

  var result resultMap
  err = getRequest(API_TICKER, &result)
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

func (client Client) LastPrice() (price float64, err error) {
  ticker, err := client.Ticker()
  if err != nil {
    return
  }
  price = ticker.Last
  return
}
