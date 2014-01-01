package main

import (
  "fmt"
)

func buy() {
  err := client.Buy(flags.price, flags.amount)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
}
