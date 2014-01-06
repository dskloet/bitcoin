package bitfinex

import (
  "bitcoin"
  "crypto/hmac"
  "crypto/sha512"
  "crypto/tls"
  "encoding/base64"
  "encoding/hex"
  "encoding/json"
  "errors"
  "fmt"
  "net/http"
  "net/url"
  "strings"
  "time"
)

type Client struct {
  apiKey    string
  apiSecret string
  nonce     int64

  currencyPair string
  http         *http.Client
}

func NewClient(apiKey, apiSecret string, insecureSkipVerify bool) *Client {
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
  }
  httpClient := &http.Client{Transport: tr}

  return &Client{
    apiKey:    apiKey,
    apiSecret: apiSecret,
    nonce:     time.Now().UnixNano() / 1000000,

    currencyPair: "btcusd",
    http:         httpClient,
  }
}

func (client *Client) createParams() (params map[string]string) {
  nonce := fmt.Sprintf("%v", client.nonce)
  client.nonce++
  params = make(map[string]string)
  params["nonce"] = nonce
  return
}

func (client Client) postRequest(
  path string, params map[string]string, result interface{}) (err error) {

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

  resp, err := client.http.Do(req)
  if client.hasError(err) {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  return
}

func (client *Client) getRequest(path string, result interface{}) (err error) {
  resp, err := client.http.Get(API_URL + path + client.currencyPair)
  if client.hasError(err) {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  return
}

func (client *Client) hasError(err error) bool {
  if err != nil {
    if _, ok := err.(*url.Error); ok {
      fmt.Printf("MacOSX may have problems with SSL certificates. " +
        "Try disabling certificate verification at your own risk.\n")
    }
  }
  return err != nil
}

func (client *Client) SetDryRun(dryRun bool) {
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

func (client *Client) Buy(price, amount float64) (err error) {
  err = errors.New("Not implemented")
  return
}

func (client *Client) Sell(price, amount float64) (err error) {
  err = errors.New("Not implemented")
  return
}

func (client *Client) CancelOrder(id bitcoin.OrderId) (err error) {
  err = errors.New("Not implemented")
  return
}

func (client *Client) UserTransactions() (
  transactions []bitcoin.UserTransaction, err error) {
  err = errors.New("Not implemented")
  return
}
