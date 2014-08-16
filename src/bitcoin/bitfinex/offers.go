package bitfinex

import (
  "fmt"
)

func (client *Client) OpenOffers() (err error) {
  var resp []offerResponse
  params := client.createParams()
  err = client.postRequest("credits", params, &resp)
  if err != nil {
    fmt.Printf("err = %v\n", err)
    return
  }

  fmt.Printf("resp = %v\n", resp)
  return
}
