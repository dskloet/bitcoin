package main

import (
  "bitstamp"
  "fmt"
)

func main() {
  client := bitstamp.NewClient(
    "86529",
    "ApdUHqBY8xTqf2Wf9xcVvTydrRdrBmWS",
    "AjGI8l3KlYyFmpZTUHWZQT7bgVTLhU3Z")

  balance, err := client.RequestBalance()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("balance = %v\n", balance)
}
