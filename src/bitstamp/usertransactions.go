package bitstamp

import (
  "strconv"
  "time"
)

type TransactionType int

type UserTransaction struct {
  Datetime time.Time
  Id       int64
  Type     TransactionType
  Usd      float64
  Btc      float64
  Fee      float64
  OrderId  int64
}

type unparsedUserTransaction struct {
  Datetime string
  Id       int64
  Type     TransactionType
  Usd      string
  Btc      string
  Fee      string
  Order_id int64
}

func (client Client) UserTransactions() (
  transactions []UserTransaction, err error) {

  params := client.createParams()
  resp, err := postRequest(API_USER_TRANSACTIONS, params)
  if err != nil {
    return
  }
  var unparsed []unparsedUserTransaction
  err = jsonParse(resp.Body, &unparsed)
  if err != nil {
    return
  }

  n := len(unparsed)
  transactions = make([]UserTransaction, n)
  for i, unparsedTx := range unparsed {
    var transaction UserTransaction
    transaction, err = parseUserTransaction(unparsedTx)
    if err != nil {
      return
    }
    transactions[n-1-i] = transaction
  }
  return
}

func parseUserTransaction(
  unparsed unparsedUserTransaction) (transaction UserTransaction, err error) {

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
  transaction = UserTransaction{
    Datetime: datetime,
    Id:       unparsed.Id,
    Type:     unparsed.Type,
    Usd:      usd,
    Btc:      btc,
    Fee:      fee,
    OrderId:  unparsed.Order_id,
  }
  return
}
