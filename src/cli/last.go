package main

import (
  "fmt"
)

func last() {
  price, err := client.LastPrice()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("%v\n", price)
}
