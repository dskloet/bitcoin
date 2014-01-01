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
  id        bitcoin.OrderId
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
  flag.StringVar(
    (*string)(&flags.id), "id", "", "Order ID for cancel_order command")
  flag.Parse()
  return
}
