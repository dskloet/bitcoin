package bitstamp

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "strconv"
)

type unparsedOrderBook struct {
  Timestamp string
  Bids      [][]string
  Asks      [][]string
}

func (client Client) OrderBook() (
  bids []bitcoin.Order, asks []bitcoin.Order, err error) {

  var unparsed unparsedOrderBook
  err = getRequest(API_ORDER_BOOK, &unparsed)
  if err != nil {
    return
  }

  for _, unparsedBid := range unparsed.Bids {
    var price, amount float64
    price, amount, err = parseFloatPair(unparsedBid)
    if err != nil {
      return
    }
    bids = append(bids, bitcoin.BuyOrder(price, amount))
  }
  for _, unparsedAsk := range unparsed.Asks {
    var price, amount float64
    price, amount, err = parseFloatPair(unparsedAsk)
    if err != nil {
      return
    }
    asks = append(asks, bitcoin.SellOrder(price, amount))
  }
  return
}

func parseFloatPair(pair []string) (a, b float64, err error) {
  a, err = strconv.ParseFloat(pair[0], 64)
  if err != nil {
    return
  }
  b, err = strconv.ParseFloat(pair[1], 64)
  return
}
