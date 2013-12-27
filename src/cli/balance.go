package main

import (
  "bitstamp"
  "fmt"
)

func balance() {
  client := bitstamp.NewClient(
      flags.clientId,
      flags.apiKey,
      flags.apiSecret)

  balance, err := client.Balance()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("USD:          %8.2f\n", balance.Usd)
  fmt.Printf("USD reserved: %8.2f\n", balance.UsdReserved)
  fmt.Printf("USD available:%8.2f\n", balance.UsdAvailable)
  fmt.Printf("BTC:           %f\n", balance.Btc)
  fmt.Printf("BTC reserved:  %f\n", balance.BtcReserved)
  fmt.Printf("BTC available: %f\n", balance.BtcAvailable)
  fmt.Printf("Fee: %.2f%%\n", balance.Fee)
}
