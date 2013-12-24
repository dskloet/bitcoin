package main

import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/json"
  "flag"
	"fmt"
  "math"
  "net/http"
  "net/url"
  "os"
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

    USD_BALANCE = "usd_balance"
    BTC_BALANCE = "btc_balance"
    FEE = "fee"
)

var flagTest bool
var flagClientId string
var flagApiKey string
var flagApiSecret string
var flagSpread float64
var flagBtcRatio float64

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

func computeBuyOrder(A, b, R, F, s float64) (amount, price float64) {
  previousRate := R*A / b
  lowRate := previousRate / s
  lowX := feeRound((R * A - b * lowRate) / (1 + R + R * F), F)
  lowRate = (((A - lowX * (1 + F)) * R) - lowX) / b
  buy := lowX / lowRate
  return buy, lowRate
}

func placeBuyOrders(A, b, R, F, s float64) (err error) {
  amount, price := computeBuyOrder(A, b, R, F, s)
  err = requestBuyOrder(amount, price)
  if err != nil {
    return
  }
  cost := amount * price * (1 + F)
  A -= cost
  b += amount
  amount, price = computeBuyOrder(A, b, R, F, s)
  err = requestBuyOrder(amount, price)
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

func computeSellOrder(A, b, R, F, s float64) (amount, price float64) {
  previousRate := R*A / b
  highRate := previousRate * s
  highX := feeRound((b * highRate - R * A) / (1 + R + R * F) * (1 + F), F)
  highRate = (((A + highX * (1 - F)) * R) + highX) / b
  sell := highX / highRate
  return sell, highRate
}

func placeSellOrders(A, b, R, F, s float64) (err error) {
  amount, price := computeSellOrder(A, b, R, F, s)
  err = requestSellOrder(amount, price)
  if err != nil {
    return
  }
  gain := amount * price * (1 - F)
  A += gain
  b -= amount
  amount, price = computeSellOrder(A, b, R, F, s)
  err = requestSellOrder(amount, price)
  return
}

func (result ApiResult) get(name string) float64 {
  value, _ := strconv.ParseFloat(result[name].(string), 64)
  return value
}

func feeRound(x, feeRate float64) float64 {
  fee := math.Ceil(x * feeRate * 100)
  return fee / (feeRate * 100)
}

func initFlags() {
  flag.BoolVar(&flagTest, "test", false, "Don't change any orders. Just output.")
  flag.StringVar(&flagApiKey, "api_key", "", "Bitstamp API key")
  flag.StringVar(&flagApiSecret, "api_secret", "", "Bitstamp API secret")
  flag.StringVar(&flagClientId, "client_id", "", "Bitstamp client ID")
  flag.Float64Var(
      &flagSpread, "spread", 2.0, "Percentage distance between buy/sell price")
  flag.Float64Var(
      &flagBtcRatio, "btc_ratio", 0.2, "Ratio of capital that should be BTC")
  flag.Parse()

  if flagApiKey == "" || flagApiSecret == "" || flagClientId == "" {
    fmt.Printf("--api_key, --api_secret, --client_id must all be specified\n")
    os.Exit(1)
  }
}

func main() {
  initFlags()

  openOrders, err := requestOrders()
  if err != nil {
    fmt.Printf("Error open orders: %v\n", err)
    return
  }
  if flagTest {
    fmt.Printf("%v open orders\n", len(openOrders))
  } else {
    if len(openOrders) == 4 {
      return
    }
  }
  if len(openOrders) != 4 {
    for _, order := range openOrders {
      err = cancelOrder(order)
      if err != nil {
        fmt.Printf("Error cancel order: %v\n", err)
        return
      }
    }
  }

  balance, err := requestMap(API_BALANCE)
  if err != nil {
    fmt.Printf("Error balance: %v\n", err)
    return
  }
  A := balance.get(USD_BALANCE)
  b := balance.get(BTC_BALANCE)
  R := flagBtcRatio / (1 - flagBtcRatio)
  F := balance.get(FEE) / 100
  if flagSpread < 200 * F {
    fmt.Printf(
        "spread (%.2f%%) must be at least twice the fee (%.2f%%) " +
        "not to make a loss.\n", flagSpread, 100 * F)
    return
  }
  s := 1 + (flagSpread / 100)

  previousRate := R*A / b

  fmt.Printf("Creating new bitstamp orders.\n")
  fmt.Printf("USD = %v\n", A)
  fmt.Printf("BTC = %v\n", b)
  fmt.Printf("Fee = %v\n", F)
  fmt.Printf("Rate = %.2f\n", previousRate)

  err = placeBuyOrders(A, b, R, F, s)
  if err != nil {
    fmt.Printf("Error buy: %v\n", err)
    return
  }

  err = placeSellOrders(A, b, R, F, s)
  if err != nil {
    fmt.Printf("Error sell: %v\n", err)
    return
  }
}
