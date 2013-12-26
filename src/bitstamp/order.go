package bitstamp

import (
  "fmt"
)

type OrderType int

type Order struct {
  Id     int64
  Type   OrderType
  Price  float64
  Amount float64
}

func (order *Order) Verb() string {
  if order.Type == ORDER_BUY {
    return "Buy"
  }
  return "Sell"
}

func (order *Order) String() string {
  return fmt.Sprintf(
    "%v %.8f at %.2f for %.2f",
    order.Verb(),
    order.Amount,
    order.Price,
    order.Amount*order.Price)
}
