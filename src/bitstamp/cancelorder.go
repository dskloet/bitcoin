package bitstamp

import (
  "errors"
  "fmt"
)

func (client *Client) CancelOrder(order Order) (err error) {
  if client.DryRun {
    fmt.Printf("Skipping cancel order %v\n", order)
    return
  }
  fmt.Printf("Cancel order %v\n", order)
  err = client.CancelOrderById(order.Id)
  return
}

func (client *Client) CancelOrderById(id int64) (err error) {
  if client.DryRun {
    fmt.Printf("Skipping cancel order %v\n", id)
    return
  }
  params := client.createParams()
  params["id"] = []string{fmt.Sprintf("%d", id)}
  resp, err := postRequest(API_CANCEL_ORDER, params)
  if err != nil {
    return
  }
  result, err := readerToString(resp.Body)
  if err != nil {
    return
  }
  if result != "true" {
    return errors.New(result)
  }
  return
}
