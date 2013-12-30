package main

import (
  "flag"
  "fmt"
  "os"
)

var flagTest bool
var flagClientId string
var flagApiKey string
var flagApiSecret string
var flagSpread float64
var flagBtcRatio float64

func initFlags() {
  flag.BoolVar(&flagTest, "test", false, "Don't change any orders. Just output.")
  flag.StringVar(&flagApiKey, "api_key", "", "Bitstamp API key")
  flag.StringVar(&flagApiSecret, "api_secret", "", "Bitstamp API secret")
  flag.StringVar(&flagClientId, "client_id", "", "Bitstamp client ID")
  flag.Float64Var(
    &flagSpread, "spread", 2.0, "Percentage distance between buy/sell price")
  flag.Float64Var(
    &flagBtcRatio, "btc_ratio", 0.5, "Ratio of capital that should be BTC")
  flag.Parse()

  if flagApiKey == "" || flagApiSecret == "" || flagClientId == "" {
    fmt.Printf("--api_key, --api_secret, --client_id must all be specified\n")
    os.Exit(1)
  }
}
