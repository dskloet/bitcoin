package btce

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "errors"
  "time"
)

type tradeHistoryResponse struct {
  Success int
  Error   string
  Return  map[string]userTrade
}

type userTrade struct {
  Type      string
  Timestamp int64
  Rate      float64
  Amount    float64
}

func (client *Client) UserTransactions() (
  transactions []bitcoin.UserTransaction, err error) {
  params := client.createParams()
  params.Set("pair", "btc_usd")
  params.Set("order", "ASC")
  var resp tradeHistoryResponse
  err = client.postRequest(API_TRADE_HISTORY, params, &resp)
  if err != nil {
    return
  }
  if resp.Success != 1 {
    err = errors.New(resp.Error)
    return
  }
  for _, trade := range resp.Return {
    datetime := time.Unix(trade.Timestamp, 0)
    fee, _ := client.Fee()
    amount := trade.Amount
    if trade.Type == "sell" {
      amount = -amount
    }
    transactions = append(transactions, bitcoin.UserTransaction{
      Datetime:       datetime,
      CurrencyAmount: -trade.Rate * amount,
      BtcAmount:      amount,
      Fee:            trade.Rate * trade.Amount * fee,
    })
  }
  return
}
