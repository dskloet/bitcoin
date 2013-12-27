package bitstamp

import (
  "encoding/json"
)

func (client *Client) OpenOrders() (openOrders []*Order, err error) {
  params := client.createParams()
  resp, err := postRequest(API_OPEN_ORDERS, params)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  var mapList []resultMap
  jsonDecoder := json.NewDecoder(resp.Body)
  err = jsonDecoder.Decode(&mapList)
  if err != nil {
    return
  }

  openOrders = make([]*Order, len(mapList))
  for i, orderMap := range mapList {
    openOrders[i] = mapToOrder(orderMap)
  }

  return
}

func mapToOrder(result resultMap) *Order {
  return &Order{
    Id:     result.getInt("id"),
    Type:   OrderType(result.getInt("type")),
    Price:  result.getFloat("price"),
    Amount: result.getFloat("amount"),
  }
}
