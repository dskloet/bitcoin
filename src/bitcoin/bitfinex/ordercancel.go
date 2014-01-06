package bitfinex

import (
  "bitcoin"
  "errors"
  "strconv"
)

func (client *Client) CancelOrder(id bitcoin.OrderId) (err error) {
  params := client.createParams()
  params["order_id"], err = strconv.ParseInt(string(id), 10, 64)
  if err != nil {
    return
  }
  var resp map[string]interface{}
  err = client.postRequest(API_ORDER_CANCEL, params, &resp)
  if err != nil {
    return
  }
  if message, ok := resp["message"]; ok {
    err = errors.New(message.(string))
  }
  return
}
