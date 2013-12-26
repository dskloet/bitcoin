package bitstamp

import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/json"
  "fmt"
  "net/http"
  "net/url"
  "time"
)

type Client struct {
  clientId  string
  apiKey    string
  apiSecret string
  nonce     int64
  DryRun    bool
}

func NewClient(clientId, apiKey, apiSecret string) *Client {
  return &Client{
    clientId:  clientId,
    apiKey:    apiKey,
    apiSecret: apiSecret,
    nonce:     time.Now().UnixNano() / 1000000,
  }
}

func (client *Client) createParams() (params url.Values) {
  nonce := fmt.Sprintf("%v", client.nonce)
  client.nonce++
  message := nonce + client.clientId + client.apiKey
  mac := hmac.New(sha256.New, []byte(client.apiSecret))
  mac.Write([]byte(message))

  params = make(url.Values)
  params["key"] = []string{client.apiKey}
  params["nonce"] = []string{nonce}
  params["signature"] = []string{fmt.Sprintf("%X", mac.Sum(nil))}
  return
}

func postRequest(path string, params url.Values) (
  resp *http.Response, err error) {

  var httpClient http.Client
  return httpClient.PostForm(API_URL+path, params)
}

func requestMap(path string, params url.Values) (
  result resultMap, err error) {

  resp, err := postRequest(path, params)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  jsonDecoder := json.NewDecoder(resp.Body)
  err = jsonDecoder.Decode(&result)
  return
}
