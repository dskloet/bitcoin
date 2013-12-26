package main

import (
  "bitstamp"
  "fmt"
)

var flags Flags = initFlags()

func ticker() {
  ticker, err := bitstamp.RequestTicker()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
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

func main() {
  if flags.ticker {
    ticker()
    return
  }

}
