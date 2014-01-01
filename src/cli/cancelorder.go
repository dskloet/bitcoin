package main

import (
  "bitstamp"
  "fmt"
)

func cancelOrder() {
  client := bitstamp.NewClient(
    flags.clientId,
    flags.apiKey,
    flags.apiSecret)
  err := client.CancelOrder(flags.id)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
}
