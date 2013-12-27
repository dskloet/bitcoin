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

func NewOrder(orderType OrderType, price, amount float64) *Order {
  return &Order{
    Type:   orderType,
    Price:  price,
    Amount: amount,
  }
}

func NewBuyOrder(price, amount float64) *Order {
  return NewOrder(ORDER_BUY, price, amount)
}

func NewSellOrder(price, amount float64) *Order {
  return NewOrder(ORDER_SELL, price, amount)
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
