package main

import (
  "flag"
)

type Flags struct {
  ticker bool
}

func initFlags() (flags Flags) {
  flag.BoolVar(&flags.ticker, "ticker", false, "Show ticker information")
  flag.Parse()
  return
}
