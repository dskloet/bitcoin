package main

import (
  "bitstamp"
  "fmt"
  "os"
)

const (
  COMMANDS = "{ticker, last, order_book, transactions, eur_usd, balance, " +
    "user_transactions, open_orders, cancel_order, buy, sell}"
)

var flags Flags
var client *bitstamp.Client

func main() {
  flags = initFlags()
  client = bitstamp.NewClient(
    flags.clientId,
    flags.apiKey,
    flags.apiSecret)
  client.SetDryRun(flags.dryRun)

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
  case "buy":
    buy()
  case "sell":
    sell()
  default:
    fmt.Printf("Command must be one of %s\n", COMMANDS)
    os.Exit(1)
  }
}
