package main

import (
  "fmt"
)

func userTransactions() {
  transactions, err := client.UserTransactions()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  for _, transaction := range transactions {
    fmt.Printf("%v	%12.8f	%14.8f	%9.4f	%5.2f\n",
      transaction.Datetime.Format("2006-01-02 15:04:05"),
      transaction.Amount,
      -transaction.Price*transaction.Amount,
      transaction.Price,
      transaction.Fee)
  }
}
