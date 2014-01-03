package bitcoin

import (
  "time"
)

type Transaction struct {
  Datetime time.Time
  Price    float64
  Amount   float64
}
