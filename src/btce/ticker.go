package btce

type Ticker struct {
  Last float64
}

func (client Client) LastPrice() (price float64, err error) {
  var tickerMap map[string]Ticker
  err = getRequest(API_TICKER, &tickerMap)
  if err != nil {
    return
  }
  price = tickerMap["ticker"].Last
  return
}
