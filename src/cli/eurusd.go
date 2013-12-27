package main

import (
  "bitstamp"
  "fmt"
)

func eurUsd() {
  var client bitstamp.Client
  eurUsd, err := client.EurUsd()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("buy: %7.4f\n", eurUsd.Buy)
  fmt.Printf("sell:%7.4f\n", eurUsd.Sell)
}
