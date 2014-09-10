package main

import (
  "bitcoin"
  "fmt"
)

func balance() {
  fiat, err := client.Balance(bitcoin.FIAT)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("Fiat: %.2f\n", fiat)

  btc, err := client.Balance(bitcoin.BTC)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("BTC:  %.8f\n", btc)

  fee, err := client.Fee()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("Fee:  %.2f%%\n", fee*100)
}
