package bitstamp

import (
  "fmt"
)

func (client *Client) RequestCancelOrder(order *Order) (err error) {
  if client.DryRun {
    fmt.Printf("Skipping cancel order %v\n", order)
    return
  }
  fmt.Printf("Cancel order %v\n", order)
  params := client.createParams()
  params["id"] = []string{fmt.Sprintf("%d", order.Id)}
  resp, err := postRequest(API_CANCEL_ORDER, params)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  return
}
