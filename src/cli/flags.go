package main

import (
  "bitcoin"
  "flag"
)

type Flags struct {
  c         string
  clientId  string
  apiKey    string
  apiSecret string
  id        int64
  first     bitcoin.Currency
  second    bitcoin.Currency
}

func initFlags() (flags Flags) {
  flag.StringVar(&flags.c, "c", "ticker", "Command. Any from "+COMMANDS)
  flag.StringVar(
    &flags.clientId, "client_id", "",
    "Bitstamp Client ID for authenticated requests")
  flag.StringVar(
    &flags.apiKey, "api_key", "",
    "Bitstamp API key for authenticated requests")
  flag.StringVar(
    &flags.apiSecret, "api_secret", "",
    "Bitstamp API secret for authenticated requests")
  flag.Int64Var(&flags.id, "id", 0, "Order ID for cancel_order command")
  flag.StringVar(
    (*string)(&flags.first), "first", "btc", "First of the currency pair")
  flag.StringVar(
    (*string)(&flags.second), "second", "usd", "Second of the currency pair")
  flag.Parse()
  return
}
