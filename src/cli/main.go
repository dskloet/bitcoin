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

  orders, err := client.RequestOpenOrders()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("orders = %v\n", orders)
}
