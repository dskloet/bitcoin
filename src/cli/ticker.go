package main

import (
  "fmt"
)

func ticker() {
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
  price, err := client.LastPrice()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("%v\n", price)
}
