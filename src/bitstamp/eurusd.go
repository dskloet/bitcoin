package bitstamp

type EurUsd struct {
  Buy  float64
  Sell float64
}

func (client Client) EurUsd() (eurUsd EurUsd, err error) {
  result, err := getMap(API_EUR_USD)
  if err != nil {
    return
  }
  eurUsd.Buy = result.getFloat("buy")
  eurUsd.Sell = result.getFloat("sell")
  return
}
