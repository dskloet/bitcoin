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

type UserTransactionList []UserTransaction

func (list UserTransactionList) Len() int {
  return len(list)
}

func (list UserTransactionList) Less(i, j int) bool {
  return list[i].Datetime.Before(list[j].Datetime)
}

func (list UserTransactionList) Swap(i, j int) {
  list[i], list[j] = list[j], list[i]
}
