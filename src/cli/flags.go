package main

import (
  "flag"
)

type Flags struct {
  c         string
  clientId  string
  apiKey    string
  apiSecret string
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
  flag.Parse()
  return
}
