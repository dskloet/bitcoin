package bitstamp

type Balance struct {
  Usd          float64
  Btc          float64
  UsdReserved  float64
  BtcReserved  float64
  UsdAvailable float64
  BtcAvailable float64
  Fee          float64
}

func (client *Client) RequestBalance() (balance Balance, err error) {
  params := client.createParams()
  result, err := requestMap("balance/", params)
  if err != nil {
    return
  }
  return Balance{
    Usd:          result.getFloat("usd_balance"),
    Btc:          result.getFloat("btc_balance"),
    UsdReserved:  result.getFloat("usd_reserved"),
    BtcReserved:  result.getFloat("btc_reserved"),
    UsdAvailable: result.getFloat("usd_available"),
    BtcAvailable: result.getFloat("btc_available"),
    Fee:          result.getFloat("fee"),
  }, nil
}
