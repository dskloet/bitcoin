package bitstamp

import (
  "sort"
)

func (client *Client) OpenOrders() (openOrders OrderList, err error) {
  params := client.createParams()
  resp, err := postRequest(API_OPEN_ORDERS, params)
  if err != nil {
    return
  }
  var mapList []resultMap
  err = jsonParse(resp.Body, &mapList)
  if err != nil {
    return
  }

  openOrders = make(OrderList, len(mapList))
  for i, orderMap := range mapList {
    openOrders[i] = mapToOrder(orderMap)
  }
  sort.Sort(openOrders)

  return
}

func mapToOrder(result resultMap) Order {
  return Order{
    Id:     result.getInt("id"),
    Type:   OrderType(result.getInt("type")),
    Price:  result.getFloat("price"),
    Amount: result.getFloat("amount"),
  }
}
