package main

import (
  "flag"
  "fmt"
  "os"
)

var flagTest bool
var flagExchange string
var flagClientId string
var flagApiKey string
var flagApiSecret string
var flagSpread float64
var flagBtcRatio float64
var flagFeeRound bool
var flagOffsetUsd float64
var flagOffsetBtc float64
var flagFeeAlwaysUsd bool

func initFlags() {
  flag.BoolVar(&flagTest, "test", false, "Don't change any orders. Just output.")
  flag.StringVar(&flagExchange, "exchange", "bitstamp", "Exchange to connect to.")
  flag.StringVar(&flagApiKey, "api_key", "", "Bitstamp API key")
  flag.StringVar(&flagApiSecret, "api_secret", "", "Bitstamp API secret")
  flag.StringVar(&flagClientId, "client_id", "", "Bitstamp client ID")
  flag.Float64Var(
    &flagSpread, "spread", 2.0, "Percentage distance between buy/sell price")
  flag.Float64Var(
    &flagBtcRatio, "btc_ratio", 0.5, "Ratio of capital that should be BTC")
  flag.BoolVar(
    &flagFeeRound, "fee_round", false,
    "Round order size up such that the fee is an integer number of cents.")
  flag.Float64Var(
    &flagOffsetUsd, "offset_usd", 0,
    "Offset the USD balance before determining which orders to make.")
  flag.Float64Var(
    &flagOffsetBtc, "offset_btc", 0,
    "Offset the BTC balance before determining which orders to make.")
  flag.BoolVar(
    &flagFeeAlwaysUsd, "fee_always_usd", false,
    "Whether the fee is always paid from USD. " +
    "Otherwise it's paid from BTC of BTC are bought.")
  flag.Parse()

  if flagApiKey == "" || flagApiSecret == "" {
    fmt.Printf("--api_key and --api_secret be specified\n")
    os.Exit(1)
  }
}
