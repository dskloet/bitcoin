package main

import (
  "fmt"
  "os"
)

const (
  COMMANDS = "{ticker, orderbook, transactions, eurusd}"
)

var flags Flags

func main() {
  flags = initFlags()
  switch flags.c {
  case "ticker":
    ticker()
  case "orderbook":
    orderBook()
  case "transactions":
    transactions()
  case "eurusd":
    eurUsd()
  default:
    fmt.Printf("Command must be one of %s\n", COMMANDS)
    os.Exit(1)
  }
}
