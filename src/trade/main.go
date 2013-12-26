package main

import (
  "bitstamp"
  "fmt"
  "math"
)

type OrderMap map[string]*StatusOrder

func (orderMap OrderMap) add(order *StatusOrder) {
  str := order.String()
  existing, ok := orderMap[str]
  if ok {
    existing.status = ORDER_KEEP
  } else {
    order.status = ORDER_NEW
    orderMap[str] = order
  }
}

func computeBuyOrder(A, b, R, F, s float64) (amount, price float64) {
  previousRate := R * A / b
  lowRate := previousRate / s
  lowX := feeRound((R*A-b*lowRate)/(1+R+R*F), F)
  lowRate = (((A - lowX*(1+F)) * R) - lowX) / b
  buy := lowX / lowRate
  return buy, lowRate
}

func placeBuyOrders(A, b, R, F, s float64, orderMap OrderMap) (err error) {
  amount, price := computeBuyOrder(A, b, R, F, s)
  orderMap.add(NewBuyOrder(amount, price))

  cost := amount * price * (1 + F)
  A -= cost
  b += amount
  amount, price = computeBuyOrder(A, b, R, F, s)
  orderMap.add(NewBuyOrder(amount, price))
  return
}

func computeSellOrder(A, b, R, F, s float64) (amount, price float64) {
  previousRate := R * A / b
  highRate := previousRate * s
  highX := feeRound((b*highRate-R*A)/(1+R+R*F)*(1+F), F)
  highRate = (((A + highX*(1-F)) * R) + highX) / b
  sell := highX / highRate
  return sell, highRate
}

func placeSellOrders(A, b, R, F, s float64, orderMap OrderMap) (err error) {
  amount, price := computeSellOrder(A, b, R, F, s)
  orderMap.add(NewSellOrder(amount, price))

  gain := amount * price * (1 - F)
  A += gain
  b -= amount
  amount, price = computeSellOrder(A, b, R, F, s)
  orderMap.add(NewSellOrder(amount, price))
  return
}

func feeRound(x, feeRate float64) float64 {
  fee := math.Ceil(x * feeRate * 100)
  return fee / (feeRate * 100)
}

func main() {
  initFlags()
  client := bitstamp.NewClient(flagClientId, flagApiKey, flagApiSecret)
  client.DryRun = flagTest

  openOrders, err := client.RequestOpenOrders()
  if err != nil {
    fmt.Printf("Error open orders: %v\n", err)
    return
  }
  if flagTest {
    fmt.Printf("%v open orders:\n", len(openOrders))
    for _, order := range openOrders {
      fmt.Printf("%v\n", order)
    }
  } else {
    if len(openOrders) == 4 {
      return
    }
  }
  orderMap := make(map[string]*StatusOrder)
  for _, order := range openOrders {
    orderMap[order.String()] = &StatusOrder{order, ORDER_REMOVE}
  }

  balance, err := client.RequestBalance()
  if err != nil {
    fmt.Printf("Error balance: %v\n", err)
    return
  }
  A := balance.Usd
  b := balance.Btc
  R := flagBtcRatio / (1 - flagBtcRatio)
  F := balance.Fee / 100
  if flagSpread < 200*F {
    fmt.Printf(
      "spread (%.2f%%) must be at least twice the fee (%.2f%%) "+
        "not to make a loss.\n", flagSpread, 100*F)
    return
  }
  s := 1 + (flagSpread / 100)

  previousRate := R * A / b

  fmt.Printf("Creating new bitstamp orders.\n")
  fmt.Printf("USD = %v\n", A)
  fmt.Printf("BTC = %v\n", b)
  fmt.Printf("Fee = %v\n", F)
  fmt.Printf("Rate = %.2f\n", previousRate)

  placeBuyOrders(A, b, R, F, s, orderMap)
  placeSellOrders(A, b, R, F, s, orderMap)

  for _, order := range orderMap {
    err := order.Execute(client)
    if err != nil {
      fmt.Printf("Error executing order: %v\n", order)
    }
  }
}
