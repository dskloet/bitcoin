package bitstamp

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "strconv"
  "time"
)

type unparsedTransaction struct {
  Date   string
  Tid    int64
  Price  string
  Amount string
}

func (client Client) Transactions() (
  transactions []bitcoin.Transaction, err error) {

  var unparsed []unparsedTransaction
  err = getRequest(API_TRANSACTIONS, &unparsed)
  if err != nil {
    return
  }

  n := len(unparsed)
  transactions = make([]bitcoin.Transaction, n)
  for i, unparsedTx := range unparsed {
    var transaction bitcoin.Transaction
    transaction, err = parseTransaction(unparsedTx)
    if err != nil {
      return
    }
    transactions[n-1-i] = transaction
  }
  return
}

func parseTransaction(
  unparsed unparsedTransaction) (transaction bitcoin.Transaction, err error) {

  timestamp, err := strconv.ParseInt(unparsed.Date, 10, 64)
  if err != nil {
    return
  }
  date := time.Unix(timestamp, 0)
  price, err := strconv.ParseFloat(unparsed.Price, 64)
  if err != nil {
    return
  }
  amount, err := strconv.ParseFloat(unparsed.Amount, 64)
  if err != nil {
    return
  }
  transaction = bitcoin.Transaction{
    Datetime: date,
    Price:    price,
    Amount:   amount,
  }
  return
}
