package main

import (
  "bitstamp"
  "fmt"
)

func ticker() {
  var client bitstamp.Client
  ticker, err := client.Ticker()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("last:%8.2f\n", ticker.Last)
  fmt.Printf("high:%8.2f\n", ticker.High)
  fmt.Printf("low: %8.2f\n", ticker.Low)
  fmt.Printf("bid: %8.2f\n", ticker.Bid)
  fmt.Printf("ask: %8.2f\n", ticker.Ask)
  fmt.Printf(
    "volume: %.8f BTC ($%.2f million)\n",
    ticker.Volume, ticker.Volume*ticker.Last/1000000)
}

func last() {
  var client bitstamp.Client
  rate, err := client.LastRate(flags.first, flags.second)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("%v/%v: %v\n", flags.first, flags.second, rate)
}
