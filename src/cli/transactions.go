package main

import (
  "bitstamp"
  "fmt"
  "time"
)

func transactions() {
  var client bitstamp.Client
  transactions, err := client.Transactions()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  for _, transaction := range transactions {
    fmt.Printf("%v,%8.2f,%12.8f\n",
      transaction.Date.In(time.UTC),
      transaction.Price,
      transaction.Amount)
  }
}
