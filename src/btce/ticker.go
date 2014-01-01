package btce

type Ticker struct {
  Last float64
}

func (client Client) LastPrice() (price float64, err error) {
  var tickerMap map[string]Ticker
  resp, err := getRequest(API_TICKER)
  if err != nil {
    return
  }
  err = jsonParse(resp.Body, &tickerMap)
  if err != nil {
    return
  }
  price = tickerMap["ticker"].Last
  return
}
