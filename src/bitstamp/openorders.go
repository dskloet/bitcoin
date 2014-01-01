package bitstamp

import (
  "bitcoin"
  "fmt"
  "sort"
)

func (client *Client) OpenOrders() (openOrders bitcoin.OrderList, err error) {
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

  openOrders = make(bitcoin.OrderList, len(mapList))
  for i, orderMap := range mapList {
    openOrders[i] = mapToOrder(orderMap)
  }
  sort.Sort(openOrders)

  return
}

func mapToOrder(result resultMap) (order bitcoin.Order) {
  price := result.getFloat("price")
  amount := result.getFloat("amount")
  orderType := result.getInt("type")
  if orderType == ORDER_BUY {
    order = bitcoin.BuyOrder(price, amount)
  } else {
    order = bitcoin.SellOrder(price, amount)
  }
  order.Id = bitcoin.OrderId(fmt.Sprintf("%v", result.getInt("id")))
  return
}
