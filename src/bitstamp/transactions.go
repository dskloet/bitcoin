package bitstamp

import (
  "strconv"
  "time"
)

type Transaction struct {
  Date   time.Time
  Tid    int64
  Price  float64
  Amount float64
}

type unparsedTransaction struct {
  Date   string
  Tid    int64
  Price  string
  Amount string
}

func (client Client) Transactions() (transactions []Transaction, err error) {
  resp, err := getRequest(API_TRANSACTIONS)
  if err != nil {
    return
  }
  var unparsed []unparsedTransaction
  err = jsonParse(resp.Body, &unparsed)
  if err != nil {
    return
  }

  n := len(unparsed)
  transactions = make([]Transaction, n)
  for i, unparsedTx := range unparsed {
    var transaction Transaction
    transaction, err = parseTransaction(unparsedTx)
    if err != nil {
      return
    }
    transactions[n - 1 - i] = transaction
  }
  return
}

func parseTransaction(
  unparsed unparsedTransaction) (transaction Transaction, err error) {

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
  transaction = Transaction{
    Date:   date,
    Tid:    unparsed.Tid,
    Price:  price,
    Amount: amount,
  }
  return
}
