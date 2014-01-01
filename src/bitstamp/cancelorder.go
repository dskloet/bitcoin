package bitstamp

import (
  "bitcoin"
  "errors"
  "fmt"
)

func (client *Client) CancelOrder(id bitcoin.OrderId) (err error) {
  if client.DryRun {
    fmt.Printf("Skipping cancel order %v\n", id)
    return
  } else {
    fmt.Printf("Cancel order %v\n", id)
  }
  params := client.createParams()
  params["id"] = []string{string(id)}
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
