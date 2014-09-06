package kraken

import (
  "bitcoin"
  "fmt"
)

type cancelOrderResponse struct {
  Error  []string
}

func (client *Client) CancelOrder(id bitcoin.OrderId) (err error) {
  fmt.Printf("Cancel order %v\n", id)
  if client.dryRun {
    fmt.Printf("Skipped\n")
    return
  }

  params := client.createParams()
  params.Set("txid", string(id))

  var resp cancelOrderResponse
  err = client.postRequest(API_CANCEL_ORDER, params, &resp)
  return
}
