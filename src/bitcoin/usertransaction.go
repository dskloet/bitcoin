package bitcoin

import (
  "time"
)

type UserTransaction struct {
  Datetime time.Time
  Usd      float64
  Btc      float64
  Fee      float64
}
