package btce

import (
  "bitcoin"
  "errors"
  "fmt"
)

func (client *Client) CancelOrder(id bitcoin.OrderId) (err error) {
  if client.dryRun {
    fmt.Printf("Skipping cancel order %v\n", id)
    return
  } else {
    fmt.Printf("Cancel order %v\n", id)
  }
  params := client.createParams()
  params["order_id"] = []string{string(id)}
  var resp tradeResponse
  err = client.postRequest(API_CANCEL_ORDER, params, &resp)
  if err != nil {
    return
  }
  if resp.Success != 1 {
    err = errors.New(resp.Error)
  }
  return
}

