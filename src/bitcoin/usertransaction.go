package bitcoin

import (
  "time"
)

type UserTransaction struct {
  Datetime       time.Time
  BtcAmount      float64
  CurrencyAmount float64
  Fee            float64
  Fee2           float64
}
