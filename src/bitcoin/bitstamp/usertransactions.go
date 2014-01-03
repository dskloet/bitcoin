package bitstamp

import (
  "bitcoin"
  "strconv"
  "time"
)

type TransactionType int

type unparsedUserTransaction struct {
  Datetime string
  Id       int64
  Type     TransactionType
  Usd      string
  Btc      string
  Fee      string
  Order_id int64
}

func (client *Client) UserTransactions() (
  transactions []bitcoin.UserTransaction, err error) {

  params := client.createParams()
  var unparsed []unparsedUserTransaction
  err = postRequest(API_USER_TRANSACTIONS, params, &unparsed)
  if err != nil {
    return
  }

  n := len(unparsed)
  transactions = make([]bitcoin.UserTransaction, n)
  for i, unparsedTx := range unparsed {
    var transaction bitcoin.UserTransaction
    transaction, err = parseUserTransaction(unparsedTx)
    if err != nil {
      return
    }
    transactions[n-1-i] = transaction
  }
  return
}

func parseUserTransaction(unparsed unparsedUserTransaction) (
  transaction bitcoin.UserTransaction, err error) {

  datetime, err :=
    time.ParseInLocation("2006-01-02 15:04:05", unparsed.Datetime, time.UTC)
  if err != nil {
    return
  }
  usd, err := strconv.ParseFloat(unparsed.Usd, 64)
  if err != nil {
    return
  }
  btc, err := strconv.ParseFloat(unparsed.Btc, 64)
  if err != nil {
    return
  }
  fee, err := strconv.ParseFloat(unparsed.Fee, 64)
  if err != nil {
    return
  }
  transaction = bitcoin.UserTransaction{
    Datetime: datetime,
    Price:    -usd / btc,
    Amount:   btc,
    Fee:      fee,
  }
  return
}
