package main

import (
  "github.com/dskloet/bitcoin/src/bitcoin/bitfinex"
  "fmt"
)

func offer() {
  status, err := client.(*bitfinex.Client).OfferStatus(flags.id)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  fmt.Printf("Date: %v\n", status.Datetime)
  fmt.Printf("Rate: %.4f\n", status.Rate)
  fmt.Printf("Period: %v\n", status.Period)
  fmt.Printf("OriginalAmount: %.8f\n", status.OriginalAmount)
  fmt.Printf("ExecutedAmount: %.8f\n", status.ExecutedAmount)
  fmt.Printf("RemainingAmount: %.8f\n", status.RemainingAmount)
}
