package main

import (
  "bitstamp"
  "fmt"
)

func userTransactions() {
  client := bitstamp.NewClient(
    flags.clientId,
    flags.apiKey,
    flags.apiSecret)
  transactions, err := client.UserTransactions()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  for _, transaction := range transactions {
    fmt.Printf("%v	%v	%12.8f	%8.2f	%7.2f	%5.2f\n",
      transaction.Type,
      transaction.Datetime.Format("2006-01-02 15:04:05"),
      transaction.Btc,
      transaction.Usd,
      -transaction.Usd/transaction.Btc,
      transaction.Fee)
  }
}
