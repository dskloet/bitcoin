package bitcoin

type Currency string

const (
  USD = "usd"
  BTC = "btc"
)

type OrderId string

type Order struct {
  Id                    OrderId
  Buy, Sell             Currency
  BuyAmount, SellAmount float64
}

type Client interface {

  ///// Unauthenticated requests

  // The price of the first expressed in the second currency.
  LatestRate(first, second Currency) (rate float64, err error)

  // Returns the part of the order book consisting of orders buying and selling
  // the given currencies.
  OrderBook(buy, sell Currency) (orders []Order, err error)

  ///// Authenticated requests

  Balance(currency Currency) (balance float64, err error)

  Fee() (fee float64, err error)

  // Creates a limit order to buy one currency for another at the rate of
  // sellAmount / buyAmount.
  Trade(buy, sell Currency, buyAmount, sellAmount float64) (err error)

  OpenOrders() (orders []Order, err error)

  CancelOrder(order Order) (err error)
  CancelOrderById(id string) (err error)
}
