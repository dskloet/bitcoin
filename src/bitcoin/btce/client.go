package btce

import (
  "bitcoin"
  "crypto/hmac"
  "crypto/sha512"
  "encoding/hex"
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

  dryRun    bool
  infoCache Info
}

func NewClient(apiKey, apiSecret string) *Client {
  return &Client{
    apiKey:    apiKey,
    apiSecret: apiSecret,
    nonce:     time.Now().Unix(),
  }
}

func (client *Client) createParams() (params url.Values) {
  nonce := fmt.Sprintf("%v", client.nonce)
  client.nonce++
  params = make(url.Values)
  params["nonce"] = []string{nonce}
  return
}

func (client Client) postRequest(
  method string, params url.Values, result interface{}) (err error) {

  params["method"] = []string{method}
  paramString := params.Encode()

  mac := hmac.New(sha512.New, []byte(client.apiSecret))
  mac.Write([]byte(paramString))
  signature := hex.EncodeToString(mac.Sum(nil))

  req, err := http.NewRequest("POST", TAPI_URL, strings.NewReader(paramString))
  if err != nil {
    return
  }
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Set("Key", client.apiKey)
  req.Header.Set("Sign", signature)

  var httpClient http.Client
  resp, err := httpClient.Do(req)
  if err != nil {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  return
}

func getRequest(path string, result interface{}) (err error) {
  resp, err := http.Get(API_URL + path)
  if err != nil {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
  return
}

func (client *Client) SetDryRun(dryRun bool) {
  client.dryRun = dryRun
}
