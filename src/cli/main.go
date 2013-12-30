package main

import (
  "fmt"
  "os"
)

const (
  COMMANDS = "{ticker, last, order_book, transactions, eur_usd, balance, " +
    "user_transactions, open_orders, cancel_order}"
)

var flags Flags

func main() {
  flags = initFlags()
  switch flags.c {
  case "ticker":
    ticker()
  case "last":
    last()
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
  case "cancel_order":
    cancelOrder()
  default:
    fmt.Printf("Command must be one of %s\n", COMMANDS)
    os.Exit(1)
  }
}
