package main

import (
  "bitstamp"
  "fmt"
)

func openOrders() {
  client := bitstamp.NewClient(
    flags.clientId,
    flags.apiKey,
    flags.apiSecret)
  orders, err := client.OpenOrders()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  for _, order := range orders {
    fmt.Printf("%v\n", order)
  }
}
