package btce

import (
  "bitcoin"
  "errors"
)

type activeOrdersResponse struct {
  Success int
  Error   string
  Return  map[bitcoin.OrderId]activeOrderResponse
}

type activeOrderResponse struct {
  Type   string
  Rate   float64
  Amount float64
}

func (client *Client) OpenOrders() (orders bitcoin.OrderList, err error) {
  params := client.createParams()
  params["pair"] = []string{"btc_usd"}
  var resp activeOrdersResponse
  err = client.postRequest(API_ACTIVE_ORDERS, params, &resp)
  if err != nil {
    return
  }
  if resp.Success != 1 {
    err = errors.New(resp.Error)
    return
  }
  for id, activeOrder := range resp.Return {
    var order bitcoin.Order
    if activeOrder.Type == "buy" {
      order = bitcoin.BuyOrder(activeOrder.Rate, activeOrder.Amount)
    } else {
      order = bitcoin.SellOrder(activeOrder.Rate, activeOrder.Amount)
    }
    order.Id = id
  }
  return
}
