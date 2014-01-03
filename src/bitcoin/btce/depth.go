package btce

import (
  "bitcoin"
)

type orderBook struct {
  Asks [][]float64
  Bids [][]float64
}

func (client Client) OrderBook() (
  bids []bitcoin.Order, asks []bitcoin.Order, err error) {

  var book orderBook
  err = getRequest(API_DEPTH, &book)
  if err != nil {
    return
  }
  for _, pair := range book.Bids {
    bids = append(bids, bitcoin.BuyOrder(pair[0], pair[1]))
  }
  for _, pair := range book.Asks {
    asks = append(asks, bitcoin.SellOrder(pair[0], pair[1]))
  }
  return
}
