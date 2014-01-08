package bitfinex

import (
  "bitcoin"
  "crypto/hmac"
  "crypto/sha512"
  "encoding/base64"
  "encoding/hex"
  "encoding/json"
  "errors"
  "fmt"
  "net/http"
  "strings"
  "time"
)

type Client struct {
  apiKey    string
  apiSecret string
  nonce     int64
  dryRun    bool

  currencyPair string
}

func NewClient(apiKey, apiSecret string) *Client {
  return &Client{
    apiKey:    apiKey,
    apiSecret: apiSecret,
    nonce:     time.Now().UnixNano() / 1000000,

    currencyPair: "btcusd",
  }
}

func (client *Client) createParams() (params map[string]interface{}) {
  nonce := fmt.Sprintf("%v", client.nonce)
  client.nonce++
  params = make(map[string]interface{})
  params["nonce"] = nonce
  return
}

func (client Client) postRequest(
  path string, params map[string]interface{}, result interface{}) (err error) {

  params["request"] = "/v1/" + path
  paramJson, err := json.Marshal(params)
  if err != nil {
    return
  }
  payload := base64.StdEncoding.EncodeToString(paramJson)

  mac := hmac.New(sha512.New384, []byte(client.apiSecret))
  mac.Write([]byte(payload))
  signature := hex.EncodeToString(mac.Sum(nil))

  req, err := http.NewRequest("POST", API_URL+path, strings.NewReader(""))
  if err != nil {
    return
  }
  req.Header.Set("X-BFX-APIKEY", client.apiKey)
  req.Header.Set("X-BFX-PAYLOAD", payload)
  req.Header.Set("X-BFX-SIGNATURE", signature)

  var httpClient http.Client
  resp, err := httpClient.Do(req)
  if err != nil {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  return
}

func (client *Client) getRequest(path string, result interface{}) (err error) {
  resp, err := http.Get(API_URL + path + client.currencyPair)
  if err != nil {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  return
}

func (client *Client) SetDryRun(dryRun bool) {
  client.dryRun = dryRun
}

func (client Client) OrderBook() (
  bids []bitcoin.Order, asks []bitcoin.Order, err error) {
  err = errors.New("Not implemented")
  return
}

func (client Client) Transactions() (
  transactions []bitcoin.Transaction, err error) {
  err = errors.New("Not implemented")
  return
}

func (client *Client) Fee() (fee float64, err error) {
  // Fee is not (yet?) available through API on Bitfinex.
  fee = 0.0012
  return
}

func (client *Client) UserTransactions() (
  transactions []bitcoin.UserTransaction, err error) {
  err = errors.New("Not implemented")
  return
}
