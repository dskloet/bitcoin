package main

import (
  "fmt"
)

func sell() {
  err := client.Sell(flags.price, flags.amount)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
}
