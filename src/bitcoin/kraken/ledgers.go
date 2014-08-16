package kraken

import (
  "bitcoin"
  "sort"
  "strconv"
  "time"
)

type ledgerInfo struct {
  Refid string
  Time float64
  Asset string
  Amount string
  Fee string
}

type ledgerResult struct {
  Ledger map[string]ledgerInfo
  Count string
}

type ledgersResponse struct {
  Error  []string
  Result ledgerResult
}

type transactionMap map[string]*bitcoin.UserTransaction

func (client *Client) UserTransactions() (
    transactions []bitcoin.UserTransaction, err error) {

  params := client.createParams()
  var resp ledgersResponse
  err = client.postRequest(API_LEDGERS, params, &resp)
  if err != nil {
    return
  }

  txMap := make(transactionMap)
  for _, info := range(resp.Result.Ledger) {
    transaction := txMap.get(info.Refid)

    transaction.Datetime = time.Unix(int64(info.Time), 0)
    var amount, fee float64
    amount, err = strconv.ParseFloat(info.Amount, 64)
    if err != nil {
      return
    }
    fee, err = strconv.ParseFloat(info.Fee, 64)
    if err != nil {
      return
    }

    if info.Asset == "XXBT" {
      transaction.BtcAmount = amount
      transaction.Fee2 = fee
    } else {
      transaction.CurrencyAmount = amount
      transaction.Fee = fee
    }
  }

  for _, tx := range txMap {
    transactions = append(transactions, *tx)
  }
  sort.Sort(bitcoin.UserTransactionList(transactions))
  return
}

func (m *transactionMap) get(refId string) (tx *bitcoin.UserTransaction) {
  tx, ok := (*m)[refId]
  if !ok {
    tx = &bitcoin.UserTransaction{}
    (*m)[refId] = tx
  }
  return tx
}
