package kraken

import (
  //"fmt"
  "strconv"
)

type feeInfo struct {
  Fee string
  MinFee string
  MaxFee string
  NextFee string
  NextVolume string
  TierVolume string
}

type tradeVolumeResult struct {
  Currency string
  Volumne string
  Fees map[string]feeInfo
}

type tradeVolumeResponse struct {
  Error  []string
  Result tradeVolumeResult
}

func (client *Client) Fee() (fee float64, err error) {
  var resp tradeVolumeResponse
  params := client.createParams()
  params.Set("pair", client.currencyPair)
  err = client.postRequest(API_TRADE_VOLUME, params, &resp)
  if err != nil {
    return
  }
  feePercent, err := strconv.ParseFloat(resp.Result.Fees[client.currencyPair].Fee, 64)
  fee = feePercent / 100.0
  return
}
