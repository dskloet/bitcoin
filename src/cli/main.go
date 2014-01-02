package main

import (
  "bitcoin"
  "bitstamp"
  "btce"
  "fmt"
  "os"
)

const (
  COMMANDS = "{last, order_book, transactions, balance, " +
    "user_transactions, open_orders, cancel_order, buy, sell}"
)

var flags Flags
var client bitcoin.Client

func main() {
  flags = initFlags()

  if flags.exchange == "bitstamp" {
    client = bitstamp.NewClient(
      flags.clientId,
      flags.apiKey,
      flags.apiSecret)
  } else if flags.exchange == "btce" {
    client = btce.NewClient(
      flags.apiKey,
      flags.apiSecret)
  }
  client.SetDryRun(flags.dryRun)

  switch flags.c {
  case "last":
    last()
  case "order_book":
    orderBook()
  case "transactions":
    transactions()
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
