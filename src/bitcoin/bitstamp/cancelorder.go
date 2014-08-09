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
  params.Set("id", string(id))
  return request(API_CANCEL_ORDER, params)
}
