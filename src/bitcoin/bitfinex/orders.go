package bitfinex

import (
  "bitcoin"
  "fmt"
  "strconv"
)

type ordersResponse []orderResponse

type orderResponse struct {
  Id               int64
  Price            string
  Remaining_amount string
  Side             string
}

func (client *Client) OpenOrders() (orders bitcoin.OrderList, err error) {
  var resp ordersResponse
  params := client.createParams()
  err = client.postRequest(API_ORDERS, params, &resp)
  if err != nil {
    return
  }

  for _, responseOrder := range resp {
    var order bitcoin.Order
    var price, amount float64
    price, err = strconv.ParseFloat(responseOrder.Price, 64)
    if err != nil {
      return
    }
    amount, err = strconv.ParseFloat(responseOrder.Remaining_amount, 64)
    if err != nil {
      return
    }
    switch responseOrder.Side {
    case "buy":
      order = bitcoin.BuyOrder(price, amount)
    case "sell":
      order = bitcoin.SellOrder(price, amount)
    }
    order.Id = bitcoin.OrderId(fmt.Sprintf("%d", responseOrder.Id))
    orders = append(orders, order)
  }
  return
}
