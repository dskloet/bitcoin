package main

import (
  "fmt"
  "os"
)

const (
  COMMANDS = "{ticker, orderbook}"
)

var flags Flags

func main() {
  flags = initFlags()
  switch flags.c {
  case "ticker":
    ticker()
  case "orderbook":
    orderBook()
  default:
    fmt.Printf("Command must be one of %s\n", COMMANDS)
    os.Exit(1)
  }
}
