package btce

import (
  "bitcoin"
  "errors"
  "time"
)

type tradeHistoryResponse struct {
  Success int
  Error string
  Return map[string]userTrade
}

type userTrade struct {
  Timestamp int64
  Rate float64
  Amount float64
}

func (client *Client) UserTransactions() (
  transactions []bitcoin.UserTransaction, err error) {
  params := client.createParams()
  params["pair"] = []string{"btc_usd"}
  params["order"] = []string{"ASC"}
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
    transactions = append(transactions, bitcoin.UserTransaction {
      Datetime: datetime,
      Usd: trade.Rate * trade.Amount,
      Btc: trade.Amount,
      Fee: trade.Rate * trade.Amount * fee,
    })
  }
  return
}
