package bitstamp

import (
  "strconv"
  "time"
)

type OrderBook struct {
  Timestamp time.Time
  Bids      []*Order
  Asks      []*Order
}

type unparsedOrderBook struct {
  Timestamp string
  Bids      [][]string
  Asks      [][]string
}

func (client Client) OrderBook() (orderBook OrderBook, err error) {
  resp, err := getRequest(API_ORDER_BOOK)
  if err != nil {
    return
  }
  var unparsed unparsedOrderBook
  err = jsonParse(resp.Body, &unparsed)
  if err != nil {
    return
  }

  timestamp, err := strconv.ParseInt(unparsed.Timestamp, 10, 64)
  if err != nil {
    return
  }
  orderBook.Timestamp = time.Unix(timestamp, 0)

  for _, unparsedBid := range unparsed.Bids {
    var price, amount float64
    price, amount, err = parseFloatPair(unparsedBid)
    if err != nil {
      return
    }
    orderBook.Bids = append(orderBook.Bids, NewBuyOrder(price, amount))
  }
  for _, unparsedAsk := range unparsed.Asks {
    var price, amount float64
    price, amount, err = parseFloatPair(unparsedAsk)
    if err != nil {
      return
    }
    orderBook.Asks = append(orderBook.Asks, NewSellOrder(price, amount))
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
