package kraken

import (
  "bitcoin"
  "errors"
  "net/http"
)

type Client struct {
  currencyPair string
}

func NewClient() *Client {
  return &Client{currencyPair: "XXBTZUSD"}
}

func (client *Client) getRequest(path string, result interface{}) (err error) {
  resp, err := http.Get(API_URL + path)
  if err != nil {
    return
  }
  err = bitcoin.JsonParse(resp.Body, result)
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
