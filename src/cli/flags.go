package main

import (
  "flag"
)

type Flags struct {
  c string
}

func initFlags() (flags Flags) {
  flag.StringVar(&flags.c, "c", "ticker", "Command. Any from "+COMMANDS)
  flag.Parse()
  return
}
