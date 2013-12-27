package bitstamp

type Ticker struct {
  Last   float64
  High   float64
  Low    float64
  Volume float64
  Bid    float64
  Ask    float64
}

func (client Client) Ticker() (ticker Ticker, err error) {
  result, err := getMap(API_TICKER)
  if err != nil {
    return
  }
  return Ticker{
    Last:   result.getFloat("last"),
    High:   result.getFloat("high"),
    Low:    result.getFloat("low"),
    Volume: result.getFloat("volume"),
    Bid:    result.getFloat("bid"),
    Ask:    result.getFloat("ask"),
  }, nil
}
