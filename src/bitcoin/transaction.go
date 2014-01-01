package bitcoin

import (
  "time"
)

type Transaction struct {
  Date   time.Time
  Price  float64
  Amount float64
}
