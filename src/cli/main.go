package main

import (
  "bitcoin"
  "bitcoin/bitfinex"
  "bitcoin/bitstamp"
  "bitcoin/btce"
  "bitcoin/kraken"
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
  } else if flags.exchange == "bitfinex" {
    client = bitfinex.NewClient(
      flags.apiKey,
      flags.apiSecret)
  } else if flags.exchange == "kraken" {
    client = kraken.NewClient(
      flags.apiKey,
      flags.apiSecret)
  }
  client.SetDryRun(flags.dryRun)

  switch flags.c {
  case "last":
    last()
  case "ob": fallthrough
  case "order_book":
    orderBook()
  case "t": fallthrough
  case "transactions":
    transactions()
  case "b": fallthrough
  case "balance":
    balance()
  case "ut": fallthrough
  case "user_transactions":
    userTransactions()
  case "oo": fallthrough
  case "open_orders":
    openOrders()
  case "cancel": fallthrough
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
