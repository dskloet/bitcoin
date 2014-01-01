package bitstamp

import (
  "bytes"
  "crypto/hmac"
  "crypto/sha256"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "net/http"
  "net/url"
  "time"
)

type Client struct {
  clientId  string
  apiKey    string
  apiSecret string
  nonce     int64
  dryRun    bool

  tickerCache  Ticker
  balanceCache Balance
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

func (client *Client) SetDryRun(dryRun bool) {
  client.dryRun = dryRun
}

func getRequest(path string) (resp *http.Response, err error) {
  var httpClient http.Client
  return httpClient.Get(API_URL + path)
}

func getMap(path string) (result resultMap, err error) {
  resp, err := getRequest(path)
  if err != nil {
    return
  }
  err = jsonParse(resp.Body, &result)
  return
}

func postRequest(path string, params url.Values) (resp *http.Response, err error) {
  var httpClient http.Client
  return httpClient.PostForm(API_URL+path, params)
}

func request(path string, params url.Values) (err error) {
  resp, err := postRequest(path, params)
  if err != nil {
    return
  }
  result, err := readerToString(resp.Body)
  if err != nil {
    return
  }
  if result != "true" {
    err = errors.New(result)
  }
  return
}

func requestMap(path string, params url.Values) (result resultMap, err error) {
  resp, err := postRequest(path, params)
  if err != nil {
    return
  }
  err = jsonParse(resp.Body, &result)
  errorString := result["error"]
  if errorString != nil {
    err = errors.New(errorString.(string))
  }
  return
}

func jsonParse(reader io.ReadCloser, result interface{}) (err error) {
  defer reader.Close()
  buf := bytes.NewBuffer(nil)
  _, err = io.Copy(buf, reader)
  if err != nil {
    return
  }
  err = json.Unmarshal(buf.Bytes(), result)
  if err != nil {
    err = errors.New(fmt.Sprintf("Couldn't parse json: %v", buf))
  }
  return
}

func readerToString(reader io.ReadCloser) (str string, err error) {
  defer reader.Close()
  buf := bytes.NewBuffer(nil)
  _, err = io.Copy(buf, reader)
  if err != nil {
    return
  }
  str = buf.String()
  return
}
