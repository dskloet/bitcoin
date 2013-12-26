package main

import (
  "bitstamp"
)

const (
  ORDER_REMOVE = iota
  ORDER_KEEP   = iota
  ORDER_NEW    = iota
)

type OrderStatus int

type StatusOrder struct {
  *bitstamp.Order
  status OrderStatus
}

func NewOrder(orderType bitstamp.OrderType, amount, price float64) *StatusOrder {
  return &StatusOrder{
    &bitstamp.Order{
      Type:   orderType,
      Amount: amount,
      Price:  price,
    },
    ORDER_NEW,
  }
}

func NewBuyOrder(amount, price float64) *StatusOrder {
  return NewOrder(bitstamp.ORDER_BUY, amount, price)
}

func NewSellOrder(amount, price float64) *StatusOrder {
  return NewOrder(bitstamp.ORDER_SELL, amount, price)
}

func (order StatusOrder) Execute(client *bitstamp.Client) (err error) {
  if order.status == ORDER_KEEP {
    return
  }
  if order.status == ORDER_REMOVE {
    return client.RequestCancelOrder(order.Order)
  }
  if order.Type == bitstamp.ORDER_BUY {
    return client.RequestBuy(order.Amount, order.Price)
  } else {
    return client.RequestSell(order.Amount, order.Price)
  }
}
