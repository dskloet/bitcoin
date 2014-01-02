package main

import (
  "bitcoin"
  "fmt"
)

func balance() {
  usd, err := client.Balance(bitcoin.USD)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("USD: %.2f\n", usd)

  btc, err := client.Balance(bitcoin.BTC)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("BTC: %.8f\n", btc)

  fee, err := client.Fee()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("Fee: %.2f%%\n", fee*100)
}
