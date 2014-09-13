package bitfinex

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "strconv"
  "time"
)

type offerResponse struct {
  Id int64
  Rate string
  Timestamp string
  Is_cancelled bool
  Original_amount string
  Executed_amount string
  Currency string
  Period int
  Direction string
  Is_live bool
  Remaining_amount string
}

type OfferStatus struct {
  Datetime time.Time
  Rate float64
  OriginalAmount float64
  ExecutedAmount float64
  RemainingAmount float64
  Period int
}

func (client *Client) OfferStatus(id bitcoin.OrderId) (
  status OfferStatus, err error) {

  var resp offerResponse
  params := client.createParams()
  params["offer_id"], err = strconv.ParseInt(string(id), 10, 64)
  err = client.postRequest("offer/status", params, &resp)
  if err != nil {
    return
  }

  rate, err := strconv.ParseFloat(resp.Rate, 64)
  if err != nil {
    return
  }
  timestamp, err := strconv.ParseFloat(resp.Timestamp, 64)
  if err != nil {
    return
  }
  originalAmount, err := strconv.ParseFloat(resp.Original_amount, 64)
  if err != nil {
    return
  }
  executedAmount, err := strconv.ParseFloat(resp.Executed_amount, 64)
  if err != nil {
    return
  }
  remainingAmount, err := strconv.ParseFloat(resp.Remaining_amount, 64)
  if err != nil {
    return
  }

  status = OfferStatus{
    Rate: rate / 365.,
    Datetime: time.Unix(int64(timestamp), 0),
    OriginalAmount: originalAmount,
    ExecutedAmount: executedAmount,
    RemainingAmount: remainingAmount,
    Period: resp.Period,
  }
  return
}
