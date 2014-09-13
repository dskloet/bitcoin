package kraken

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "crypto/hmac"
  "crypto/sha256"
  "crypto/sha512"
  "encoding/base64"
  "errors"
  "fmt"
  "net/http"
  "net/url"
  "reflect"
  "strings"
  "time"
)

type Client struct {
  currencyPair string
  apiKey    string
  apiSecret string
  nonce     int64
  dryRun    bool
}

func NewClient(apiKey, apiSecret string) *Client {
  return &Client{
    currencyPair: "XXBTZEUR",
    apiKey:    apiKey,
    apiSecret: apiSecret,
    nonce:     time.Now().UnixNano() / 1000,
  }
}

func (client *Client) SetDryRun(dryRun bool) {
  client.dryRun = dryRun
}

func (client *Client) getRequest(path string, result interface{}) (err error) {
  resp, err := http.Get(API_URL + path)
  if err != nil {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  return
}

func (client *Client) postRequest(
    path string, params url.Values, result interface{}) (err error) {

  client.nonce++
  nonce := fmt.Sprintf("%v", client.nonce)
  params.Set("nonce", nonce)
  data := params.Encode()

  req, err := http.NewRequest("POST", API_URL + path, strings.NewReader(data))

  secret, err := base64.StdEncoding.DecodeString(client.apiSecret)
  mac := hmac.New(sha512.New, secret)
  nonceData256 := sha256.Sum256([]byte(nonce + string(data)))

  mac.Write([]byte(path + string(nonceData256[:32])))

  signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

  req.Header.Set("API-Key", client.apiKey)
  req.Header.Set("API-Sign", signature)

  var httpClient http.Client
  resp, err := httpClient.Do(req)
  if err != nil {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  if err != nil {
    return
  }

  resultValue := reflect.ValueOf(result)
  error := (resultValue.Elem().FieldByName("Error").Interface()).([]string)
  if len(error) > 0 {
    err = errors.New(strings.Join(error, "; "))
  }

  return
}

func (client *Client) createParams() (params url.Values) {
  params = make(url.Values)
  return
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
