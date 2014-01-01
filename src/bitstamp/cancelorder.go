package bitstamp

import (
  "bitcoin"
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
  params["id"] = []string{string(id)}
  return request(API_CANCEL_ORDER, params)
}
