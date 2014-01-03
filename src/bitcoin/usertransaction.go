package bitcoin

import (
  "time"
)

type UserTransaction struct {
  Datetime time.Time
  Price    float64
  Amount   float64
  Fee      float64
}
