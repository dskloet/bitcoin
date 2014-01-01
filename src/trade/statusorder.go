package main

import (
  "bitcoin"
  "bitstamp"
  "fmt"
)

const (
  ORDER_REMOVE = iota
  ORDER_KEEP   = iota
  ORDER_NEW    = iota
)

type OrderStatus int

type StatusOrder struct {
  bitcoin.Order
  status OrderStatus
}

func NewBuyOrder(price, amount float64) *StatusOrder {
  return &StatusOrder{bitcoin.BuyOrder(price, amount), ORDER_NEW}
}

func NewSellOrder(price, amount float64) *StatusOrder {
  return &StatusOrder{bitcoin.SellOrder(price, amount), ORDER_NEW}
}

func (order StatusOrder) Execute(client *bitstamp.Client) (err error) {
  if order.status == ORDER_KEEP {
    fmt.Printf("Keep order %v\n", order)
    return
  }
  if order.status == ORDER_REMOVE {
    return client.CancelOrder(order.Order.Id)
  }
  if order.Type == bitcoin.BUY_ORDER {
    return client.Buy(order.Price, order.Amount)
  } else {
    return client.Sell(order.Price, order.Amount)
  }
}
