package main

import (
  "fmt"
)

func openOrders() {
  orders, err := client.OpenOrders()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  for _, order := range orders {
    fmt.Printf("%v: %v\n", order.Id, order)
  }
}
