package btce

import (
  "bitcoin"
  "time"
)

type transaction struct {
  Date   int64
  Price  float64
  Amount float64
}

func (client Client) Transactions() (
  transactions []bitcoin.Transaction, err error) {

  var txs []transaction
  err = getRequest(API_TRADES, &txs)
  if err != nil {
    return
  }
  n := len(txs)
  transactions = make([]bitcoin.Transaction, n)
  for i, tx := range txs {
    date := time.Unix(tx.Date, 0)
    transactions[n-1-i] = bitcoin.Transaction{
      Date:   date,
      Price:  tx.Price,
      Amount: tx.Amount,
    }
  }
  return
}
