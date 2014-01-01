package bitcoin

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

func BuyOrder(price, amount float64) Order {
  return Order{Type: BUY_ORDER, Price: price, Amount: amount}
}

func SellOrder(price, amount float64) Order {
  return Order{Type: SELL_ORDER, Price: price, Amount: amount}
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
