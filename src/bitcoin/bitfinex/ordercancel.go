package bitfinex

import (
  "github.com/dskloet/bitcoin/src/bitcoin"
  "errors"
  "fmt"
  "strconv"
)

func (client *Client) CancelOrder(id bitcoin.OrderId) (err error) {
  if client.dryRun {
    fmt.Printf("Skipping cancel order %v\n", id)
    return
  } else {
    fmt.Printf("Cancel order %v\n", id)
  }
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
