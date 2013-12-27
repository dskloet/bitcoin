package main

import (
  "bitstamp"
  "fmt"
)

func orderBook() {
  var client bitstamp.Client
  orderBook, err := client.OrderBook()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("%s%20s\n", "Bids", "Asks")
  bids := orderBookTable(orderBook.Bids)
  asks := orderBookTable(orderBook.Asks)
  for i := 0; i < len(bids) || i < len(asks); i++ {
    bid := ""
    if i < len(bids) {
      bid = bids[i]
    }
    fmt.Printf("%20s", bid)
    if i < len(asks) {
      fmt.Printf("%20s", asks[i])
    }
    fmt.Printf("\n")
  }
}

func orderBookTable(orders []*bitstamp.Order) (output []string) {
  output = append(output, "      Price:  Amount")
  thresholdSteps := 0
  threshold := 0.5
  sum := 0.0
  for _, order := range orders {
    sum += order.Amount
    if sum >= threshold {
      output = append(output, fmt.Sprintf("%8.2f:%8.1f", order.Price, sum))
      for threshold <= sum {
        thresholdSteps++
        if thresholdSteps%3 == 0 {
          threshold *= 2.5
        } else {
          threshold *= 2
        }
      }
    }
  }
  return
}
