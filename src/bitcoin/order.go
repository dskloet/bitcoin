package bitcoin

import (
  "fmt"
  "strconv"
)

type OrderId string
type OrderType bool

const (
  BUY_ORDER  = true
  SELL_ORDER = false
)

type Order struct {
  Id            OrderId
  Type          OrderType
  Price, Amount float64
}

func MakeOrder(orderType OrderType, price, amount float64) Order {
  return Order{
    Type: orderType,
    Price: round(price, 5),
    Amount: round(amount, 8),
  }
}

func round(value float64, places int) (result float64) {
  format := fmt.Sprintf("%%.%df", places)
  str := fmt.Sprintf(format, value)
  result, _ = strconv.ParseFloat(str, 64)
  return
}

func BuyOrder(price, amount float64) Order {
  return MakeOrder(BUY_ORDER, price, amount)
}

func SellOrder(price, amount float64) Order {
  return MakeOrder(SELL_ORDER, price, amount)
}

func (order Order) Cancel(client Client) error {
  return client.CancelOrder(order.Id)
}

func (order Order) Execute(client Client) error {
  if order.Type == BUY_ORDER {
    return client.Buy(order.Price, order.Amount)
  }
  return client.Sell(order.Price, order.Amount)
}

func (order Order) verb() string {
  if order.Type == BUY_ORDER {
    return "Buy"
  }
  return "Sell"
}

func (order Order) String() string {
  return fmt.Sprintf(
    "%v %.8f at %.5f for %.8f",
    order.verb(),
    order.Amount,
    order.Price,
    order.Amount*order.Price)
}

type OrderList []Order

func (list OrderList) Len() int {
  return len(list)
}

func (list OrderList) Less(i, j int) bool {
  return list[j].Price < list[i].Price
}

func (list OrderList) Swap(i, j int) {
  list[i], list[j] = list[j], list[i]
}
