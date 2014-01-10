package bitfinex

import (
  "bitcoin"
  "strconv"
)

type bookResponse struct {
  Bids []bookOrderResponse
  Asks []bookOrderResponse
}

type bookOrderResponse struct {
  Price  string
  Amount string
}

func (client Client) OrderBook() (
  bids []bitcoin.Order, asks []bitcoin.Order, err error) {

  var resp bookResponse
  err = client.getRequest(API_BOOK+client.currencyPair, &resp)
  if err != nil {
    return
  }
  err = parseOrders(resp.Bids, bitcoin.BuyOrder, &bids)
  if err != nil {
    return
  }
  err = parseOrders(resp.Asks, bitcoin.SellOrder, &asks)
  return
}

func parseOrders(
  orders []bookOrderResponse,
  factory func(price, amount float64) bitcoin.Order,
  out *[]bitcoin.Order) (err error) {

  for _, order := range orders {
    var price, amount float64
    price, err = strconv.ParseFloat(order.Price, 64)
    if err != nil {
      return
    }
    amount, err = strconv.ParseFloat(order.Amount, 64)
    if err != nil {
      return
    }
    *out = append(*out, factory(price, amount))
  }
  return
}
