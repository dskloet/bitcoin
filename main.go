package main

import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/json"
	"fmt"
  "math"
  "net/http"
  "net/url"
  "strconv"
  "time"
  "io"
  "os"
)

const (
  API_KEY = "jdAfYqOa1s2VAzlFH8MS1L7IUnk3dOLw"
  API_SECRET = "yA2Wy0JmbjUas6y9KpCukj0TdMMX7oKL"
  CLIENT_ID = "86529"
)

const (
    API_URL = "https://www.bitstamp.net/api/"
    API_BALANCE = "balance/"
    API_OPEN_ORDERS = "open_orders/"
    API_CANCEL_ORDER = "cancel_order/"
    API_BUY = "buy/"
    API_SELL = "sell/"

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
  message := nonce + CLIENT_ID + API_KEY
  mac := hmac.New(sha256.New, []byte(API_SECRET))
  mac.Write([]byte(message))

  params = make(url.Values)
  params["key"] = []string{ API_KEY }
  params["nonce"] = []string{ nonce }
  params["signature"] = []string{ fmt.Sprintf("%X", mac.Sum(nil)) }
  //fmt.Printf("params = %s\n", params);
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

  //io.Copy(os.Stdout, resp.Body)
  jsonDecoder := json.NewDecoder(resp.Body)
  err = jsonDecoder.Decode(&result)
  return
}

func cancelOrder(order Order) (err error) {
  params := createParams()
  params["id"] = []string{ fmt.Sprintf("%d", order.Id) }
  resp, err := postRequest(API_CANCEL_ORDER, params)
  io.Copy(os.Stdout, resp.Body)
  return
}

func buyOrder(amount, price float64) (err error) {
  params := createParams()
  params["amount"] = []string{ fmt.Sprintf("%.8f", amount) }
  params["price"] = []string{ fmt.Sprintf("%.2f", price) }
  //fmt.Printf("buy %v\n", params)
  _, err = postRequest(API_BUY, params)
  //io.Copy(os.Stdout, resp.Body)
  return
}

func sellOrder(amount, price float64) (err error) {
  params := createParams()
  params["amount"] = []string{ fmt.Sprintf("%.8f", amount) }
  params["price"] = []string{ fmt.Sprintf("%.2f", price) }
  //fmt.Printf("sell %v\n", params)
  _, err = postRequest(API_SELL, params)
  //io.Copy(os.Stdout, resp.Body)
  return
}

func (result ApiResult) get(name string) float64 {
  value, _ := strconv.ParseFloat(result[name].(string), 64)
  return value;
}

func feeRound(x, feeRate float64) float64 {
  fee := math.Ceil(x * feeRate * 100)
  return fee / (feeRate * 100);
}

func main() {
  openOrders, err := requestOrders()
  if err != nil {
    fmt.Printf("Error open orders: %v\n", err)
    return
  }
  if len(openOrders) == 2 {
    fmt.Printf(".")
    return
  }
  fmt.Printf("\n")

  if len(openOrders) == 1 {
    cancelOrder(openOrders[0])
    if err != nil {
      fmt.Printf("Error cancel order: %v\n", err)
      return
    }
  }

  balance, err := requestMap(API_BALANCE)
  if err != nil {
    fmt.Printf("Error balance: %v\n", err)
    return
  }
  A := balance.get(USD_BALANCE)
  b := balance.get(BTC_BALANCE)
  R := 0.25
  F := balance.get(FEE) / 100
  s := 1.03

  previousRate := R*A / b
  highRate := previousRate * s
  lowRate := previousRate / s

  lowX := feeRound((R * A - b * lowRate) / (1 + R + R * F), F)
  highX := feeRound((b * highRate - R * A) / (1 + R + R * F) * (1 + F), F)
  lowRate = (((A - lowX * (1 + F)) * R) - lowX) / b
  highRate = (((A + highX * (1 - F)) * R) + highX) / b
  buy := lowX / lowRate
  sell := highX / highRate

  fmt.Printf("A = %v\n", A)
  fmt.Printf("b = %v\n", b)
  fmt.Printf("F = %v\n", F)
  fmt.Printf("E = %v %v %v\n", lowRate, previousRate, highRate)
  fmt.Printf("Buy %.8f at %.2f for %.2f\n", buy, lowRate, lowX)
  fmt.Printf("Sell %.8f at %.2f for %.2f\n", sell, highRate, highX)

  buyOrder(buy, lowRate)
  if err != nil {
    fmt.Printf("Error buy: %v\n", err)
    return
  }
  sellOrder(sell, highRate)
  if err != nil {
    fmt.Printf("Error sell: %v\n", err)
    return
  }
}
