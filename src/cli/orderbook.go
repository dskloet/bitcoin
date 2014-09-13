package main

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "fmt"
)

func orderBook() {
  bids, asks, err := client.OrderBook()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("%s%20s\n", "Bids", "Asks")
  bidsTable := orderBookTable(bids)
  asksTable := orderBookTable(asks)
  for i := 0; i < len(bidsTable) || i < len(asksTable); i++ {
    bid := ""
    if i < len(bidsTable) {
      bid = bidsTable[i]
    }
    fmt.Printf("%20s", bid)
    if i < len(asksTable) {
      fmt.Printf("%20s", asksTable[i])
    }
    fmt.Printf("\n")
  }
}

func orderBookTable(orders []bitcoin.Order) (output []string) {
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
