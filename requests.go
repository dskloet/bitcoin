package bitstamp

import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/json"
	"fmt"
  "net/http"
  "net/url"
  "strconv"
  "time"
)

const (
    API_URL = "https://www.bitstamp.net/api/"
    API_BALANCE = "balance/"
    API_OPEN_ORDERS = "open_orders/"
    API_CANCEL_ORDER = "cancel_order/"
    API_BUY = "buy/"
    API_SELL = "sell/"
)

const (
    USD_BALANCE = "usd_balance"
    BTC_BALANCE = "btc_balance"
    FEE = "fee"
)

type ApiResult map[string]interface{}

type Order struct {
  Id int64
}

var now int64 = time.Now().Unix()

func createParams() (params url.Values) {
  nonce := fmt.Sprintf("%v", now)
  now++
  message := nonce + flagClientId + flagApiKey
  mac := hmac.New(sha256.New, []byte(flagApiSecret))
  mac.Write([]byte(message))

  params = make(url.Values)
  params["key"] = []string{ flagApiKey }
  params["nonce"] = []string{ nonce }
  params["signature"] = []string{ fmt.Sprintf("%X", mac.Sum(nil)) }
  return
}

func postRequest(path string, params url.Values) (resp *http.Response, err error) {
  var client http.Client
  return client.PostForm(API_URL + path, params)
}

func requestMap(path string) (result ApiResult, err error) {
  params := createParams()
  resp, err := postRequest(path, params)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  jsonDecoder := json.NewDecoder(resp.Body)
  jsonDecoder.Decode(&result)
  return
}

func requestOrders() (result []Order, err error) {
  params := createParams()
  resp, err := postRequest(API_OPEN_ORDERS, params)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  jsonDecoder := json.NewDecoder(resp.Body)
  err = jsonDecoder.Decode(&result)
  return
}

func cancelOrder(order Order) (err error) {
  if flagTest {
    fmt.Printf("Cancel order %v\n", order)
    fmt.Printf("Skipped\n");
    return
  }
  params := createParams()
  params["id"] = []string{ fmt.Sprintf("%d", order.Id) }
  resp, err := postRequest(API_CANCEL_ORDER, params)
  if err == nil {
    resp.Body.Close()
  }
  return
}

func requestBuyOrder(amount, price float64) (err error) {
  fmt.Printf("Buy %.8f at %.2f for %.2f\n", amount, price, amount * price)
  if flagTest {
    fmt.Printf("Skipped\n")
    return
  }
  params := createParams()
  params["amount"] = []string{ fmt.Sprintf("%.8f", amount) }
  params["price"] = []string{ fmt.Sprintf("%.2f", price) }
  _, err = postRequest(API_BUY, params)
  return
}

func requestSellOrder(amount, price float64) (err error) {
  fmt.Printf("Sell %.8f at %.2f for %.2f\n", amount, price, amount * price)
  if flagTest {
    fmt.Printf("Skipped\n")
    return
  }
  params := createParams()
  params["amount"] = []string{ fmt.Sprintf("%.8f", amount) }
  params["price"] = []string{ fmt.Sprintf("%.2f", price) }
  _, err = postRequest(API_SELL, params)
  return
}

func (result ApiResult) get(name string) float64 {
  value, _ := strconv.ParseFloat(result[name].(string), 64)
  return value
}