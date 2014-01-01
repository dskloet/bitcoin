package bitcoin

type Currency string

const (
  USD = "usd"
  BTC = "btc"
)

// A client for trading one specific currency (e.g. BTC) on one particular
// exchange with prices expressed in another specific currency (e.g. USD).
type Client interface {
  SetDryRun(dryRun bool)

  ///// Unauthenticated requests

  LastPrice() (price float64, err error)
  OrderBook() (bids []Order, asks []Order, err error)

  // Returns recent transactions. How many depends on the exchange API.
  Transactions() (transactions []Transaction, err error)

  ///// Authenticated requests

  Balance(currency Currency) (balance float64, err error)
  Fee() (fee float64, err error)

  Buy(price, amount float64) (err error)
  Sell(price, amount float64) (err error)

  OpenOrders() (orders OrderList, err error)
  CancelOrder(id OrderId) (err error)

  UserTransactions() (transactions []UserTransaction, err error)
}
