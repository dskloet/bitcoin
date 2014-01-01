package main

import (
  "fmt"
)

func cancelOrder() {
  err := client.CancelOrder(flags.id)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
}
