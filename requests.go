package main

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

const (
    ORDER_BUY = 0
    ORDER_SELL = 1
)

const (
    ORDER_REMOVE = iota
    ORDER_KEEP = iota
    ORDER_NEW = iota
)

type Order struct {
  Id int64
  Type int
  Price string
  Amount string
  status int
  value string
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

func requestOrders() (result []*Order, err error) {
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
  fmt.Printf("Cancel order %v\n", order.Desc())
  if flagTest {
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

func requestOrder(order Order) (err error) {
  fmt.Printf("%v\n", order.Desc())
  if flagTest {
    fmt.Printf("Skipped\n")
    return
  }
  params := createParams()
  params["amount"] = []string{ order.Amount }
  params["price"] = []string{ order.Price }
  if order.Type == ORDER_BUY {
    _, err = postRequest(API_BUY, params)
  } else if order.Type == ORDER_SELL {
    _, err = postRequest(API_SELL, params)
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

func NewOrder(orderType int, amount, price float64) *Order {
  return &Order{
    Type: orderType,
    Amount: fmt.Sprintf("%.8f", amount),
    Price: fmt.Sprintf("%.2f", price),
    value: fmt.Sprintf("%.2f", price * amount),
  }
}

func NewBuyOrder(amount, price float64) *Order {
  return NewOrder(ORDER_BUY, amount, price)
}

func NewSellOrder(amount, price float64) *Order {
  return NewOrder(ORDER_SELL, amount, price)
}

func (order Order) Verb() string {
  if order.Type == ORDER_BUY {
    return "Buy"
  } else {
    return "Sell"
  }
}

func (order Order) String() string {
  return fmt.Sprintf("%v:%v:%v", order.Verb(), order.Amount, order.Price)
}

func (order Order) Desc() string {
  desc := fmt.Sprintf(
      "%v %v at %v",
      order.Verb(),
      order.Amount,
      order.Price)
  if order.value != "" {
    desc += " for " + order.value
  }
  return desc
}

func (order Order) Execute() (err error) {
  if order.status == ORDER_KEEP {
    return
  }
  if order.status == ORDER_REMOVE {
    return cancelOrder(order)
  }
  return requestOrder(order)
}
