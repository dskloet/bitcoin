package bitstamp

import (
  "fmt"
)

func (client *Client) Buy(price, amount float64) (err error) {
  fmt.Printf("Buy %.8f at %.2f for %.2f\n", amount, price, amount*price)
  if client.DryRun {
    fmt.Printf("Skipped\n")
    return
  }
  params := client.createParams()
  params["amount"] = []string{fmt.Sprintf("%.8f", amount)}
  params["price"] = []string{fmt.Sprintf("%.2f", price)}

  resp, err := postRequest(API_BUY, params)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  return
}
