package bitfinex

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "strconv"
  "time"
)

type myTradeResponse struct {
  Type      string
  Price     string
  Amount    string
  Timestamp string
}

func (client *Client) UserTransactions() (
  transactions []bitcoin.UserTransaction, err error) {

  var resp []myTradeResponse
  params := client.createParams()
  params["symbol"] = client.currencyPair
  err = client.postRequest(API_MYTRADES, params, &resp)
  if err != nil {
    return
  }
  n := len(resp)
  transactions = make([]bitcoin.UserTransaction, n)
  for i, myTrade := range resp {
    var transaction bitcoin.UserTransaction
    transaction, err = parseMyTrade(myTrade)
    if err != nil {
      return
    }
    transactions[n-1-i] = transaction
  }
  return
}

func parseMyTrade(trade myTradeResponse) (
  transaction bitcoin.UserTransaction, err error) {

  price, err := strconv.ParseFloat(trade.Price, 64)
  if err != nil {
    return
  }
  transaction.BtcAmount, err = strconv.ParseFloat(trade.Amount, 64)
  if err != nil {
    return
  }
  if trade.Type == "Sell" {
    transaction.BtcAmount *= -1
  }
  transaction.CurrencyAmount = -price * transaction.BtcAmount
  timestamp, err := strconv.ParseFloat(trade.Timestamp, 64)
  if err != nil {
    return
  }
  transaction.Datetime = time.Unix(int64(timestamp), 0)
  // TODO: Add fee once it's available through the API.
  return
}
