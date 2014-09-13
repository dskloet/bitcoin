package main

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "flag"
)

type Flags struct {
  exchange  string
  c         string
  dryRun    bool
  clientId  string
  apiKey    string
  apiSecret string
  id        bitcoin.OrderId
  price     float64
  amount    float64
}

func initFlags() (flags Flags) {
  flag.StringVar(
    &flags.exchange, "exchange", "bitstamp",
    "Exchange from {bitstamp, btce, bitfinex}")
  flag.StringVar(&flags.c, "c", "ticker", "Command. Any from "+COMMANDS)
  flag.BoolVar(&flags.dryRun, "dry_run", false, "Don't make/cancel any order")
  flag.StringVar(
    &flags.clientId, "client_id", "",
    "Bitstamp Client ID for authenticated requests")
  flag.StringVar(
    &flags.apiKey, "api_key", "",
    "Bitstamp API key for authenticated requests")
  flag.StringVar(
    &flags.apiSecret, "api_secret", "",
    "Bitstamp API secret for authenticated requests")
  flag.StringVar(
    (*string)(&flags.id), "id", "", "Order ID for cancel_order command")
  flag.Float64Var(&flags.price, "price", 0, "Price for buy/sell orders")
  flag.Float64Var(&flags.amount, "amount", 0, "Amount for buy/sell orders")

  flag.Parse()
  return
}
