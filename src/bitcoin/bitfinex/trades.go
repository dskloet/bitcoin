package bitfinex

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "fmt"
  "strconv"
  "time"
)

type tradeResponse struct {
  Price     string
  Amount    string
  Exchange  string
  Timestamp int64
}

func (client Client) Transactions() (
  transactions []bitcoin.Transaction, err error) {

  var resp []tradeResponse
  err = client.getRequest(API_TRADES+client.currencyPair+
    fmt.Sprintf("?timestamp=%d&limit_trades=1000",
      time.Now().Add(-time.Hour).Unix()), &resp)
  if err != nil {
    return
  }
  n := len(resp)
  for i, _ := range resp {
    trade := resp[n-1-i]
    if trade.Exchange != "bitfinex" {
      continue
    }
    var transaction bitcoin.Transaction
    transaction, err = parseTrade(trade)
    if err != nil {
      return
    }
    transactions = append(transactions, transaction)
  }
  return
}

func parseTrade(trade tradeResponse) (transaction bitcoin.Transaction, err error) {
  transaction.Price, err = strconv.ParseFloat(trade.Price, 64)
  if err != nil {
    return
  }
  transaction.Amount, err = strconv.ParseFloat(trade.Amount, 64)
  transaction.Datetime = time.Unix(trade.Timestamp, 0)
  return
}
