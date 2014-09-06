package kraken

import (
  "bitcoin"
  "strconv"
)

type orderDescription struct {
  Pair string
  Type string
  Ordertype string
  Price string
}

type orderResult struct {
  Status string
  Descr orderDescription
  Vol string
}

type openOrdersResult struct {
  Open map[bitcoin.OrderId]orderResult
}

type openOrdersResponse struct {
  Error  []string
  Result openOrdersResult
}

func (client *Client) OpenOrders() (orders bitcoin.OrderList, err error) {
  params := client.createParams()
  var resp openOrdersResponse
  err = client.postRequest(API_OPEN_ORDERS, params, &resp)
  if err != nil {
    return
  }

  for id, order := range(resp.Result.Open) {
    if order.Status != "open" || order.Descr.Ordertype != "limit" {
      continue;
    }
    var price, amount float64
    price, err = strconv.ParseFloat(order.Descr.Price, 64)
    if err != nil {
      return
    }
    amount, err = strconv.ParseFloat(order.Vol, 64)
    if err != nil {
      return
    }
    if order.Descr.Type == "buy" {
      o := bitcoin.BuyOrder(price, amount)
      o.Id = id
      orders = append(orders, o)
    } else if order.Descr.Type == "sell" {
      o := bitcoin.SellOrder(price, amount)
      o.Id = id
      orders = append(orders, o)
    }
  }
  return
}
