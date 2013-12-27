package main

import (
  "fmt"
  "os"
)

const (
  COMMANDS = "{ticker, orderbook, transactions, eurusd}"
)

var flags Flags

func main() {
  flags = initFlags()
  switch flags.c {
  case "ticker":
    ticker()
  case "order_book":
    orderBook()
  case "transactions":
    transactions()
  case "eur_usd":
    eurUsd()
  case "balance":
    balance()
  case "user_transactions":
    userTransactions()
  case "open_orders":
    openOrders()
  default:
    fmt.Printf("Command must be one of %s\n", COMMANDS)
    os.Exit(1)
  }
}
