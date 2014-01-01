package btce

import (
  "bitcoin"
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "net/http"
)

type Client struct {
}

func NewClient() *Client {
  return &Client{}
}

func getRequest(path string) (resp *http.Response, err error) {
  var httpClient http.Client
  return httpClient.Get(API_URL + path)
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

func (client *Client) Balance(currency bitcoin.Currency) (
  balance float64, err error) {
  err = errors.New("Not implemented")
  return
}

func (client *Client) Fee() (fee float64, err error) {
  err = errors.New("Not implemented")
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

func (client *Client) OpenOrders() (orders bitcoin.OrderList, err error) {
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
