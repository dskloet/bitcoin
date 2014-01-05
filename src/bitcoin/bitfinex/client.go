package bitfinex

import (
  "bitcoin"
  "crypto/tls"
  "errors"
  "fmt"
  "net/http"
  "net/url"
)

type Client struct {
  currencyPair string
  insecureSkipVerify bool
}

func NewClient(insecureSkipVerify bool) *Client {
  return &Client{
    currencyPair: "btcusd",
    insecureSkipVerify: insecureSkipVerify,
  }
}

func (client *Client) getRequest(path string, result interface{}) (err error) {
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify : client.insecureSkipVerify},
  }
  httpClient := &http.Client{Transport: tr}
  resp, err := httpClient.Get(API_URL + path + client.currencyPair)
  if err != nil {
    if _, ok := err.(*url.Error); ok {
      fmt.Printf("MacOSX may have problems with SSL certificates. " +
          "Try disabling certificate verification at your own risk.\n")
    }
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
