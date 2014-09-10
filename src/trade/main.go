package main

import (
  "bitcoin"
  "bitcoin/bitfinex"
  "bitcoin/bitstamp"
  "bitcoin/kraken"
  "fmt"
  "math"
  "os"
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

func computeBuyOrder(A, b, R, F, s float64) (price, amount float64) {
  previousRate := R * A / b
  lowRate := previousRate / s
  var factor float64
  if flagFeeAlwaysUsd {
    factor = 1 + R + R*F
  } else {
    factor = 1 + R - F
  }
  lowX := (R*A - b*lowRate) / factor
  if lowX / lowRate < flagMinAmount {
    lowRate = R * A / (b + factor * flagMinAmount)
    lowX = lowRate * flagMinAmount
  }
  if flagFeeRound {
    lowX = feeRound(lowX, F)
    lowRate = (((A - lowX*(1+F)) * R) - lowX) / b
  }
  buy := lowX / lowRate
  price, amount = lowRate, buy
  return
}

func placeBuyOrders(A, b, R, F, s float64, orderMap OrderMap) (err error) {
  price, amount := computeBuyOrder(A, b, R, F, s)
  orderMap.add(NewBuyOrder(price, amount))

  cost := amount * price * (1 + F)
  A -= cost
  b += amount
  price, amount = computeBuyOrder(A, b, R, F, s)
  orderMap.add(NewBuyOrder(price, amount))
  return
}

func computeSellOrder(A, b, R, F, s float64) (price, amount float64) {
  previousRate := R * A / b
  highRate := previousRate * s
  highX := (b*highRate - R*A) / (1 + R - R * F)
  if highX / highRate < flagMinAmount {
    highRate = R * A / (b - (1 + R - R * F) * flagMinAmount)
    highX = highRate * flagMinAmount
  }
  if flagFeeRound {
    highX = feeRound(highX, F)
    highRate = (((A + highX*(1-F)) * R) + highX) / b
  }
  sell := highX / highRate
  price, amount = highRate, sell
  return
}

func placeSellOrders(A, b, R, F, s float64, orderMap OrderMap) (err error) {
  price, amount := computeSellOrder(A, b, R, F, s)
  orderMap.add(NewSellOrder(price, amount))

  gain := amount * price * (1 - F)
  A += gain
  b -= amount
  price, amount = computeSellOrder(A, b, R, F, s)
  orderMap.add(NewSellOrder(price, amount))
  return
}

func feeRound(x, feeRate float64) float64 {
  fee := math.Ceil(x * feeRate * 100)
  return fee / (feeRate * 100)
}

func main() {
  initFlags()
  var client bitcoin.Client
  if flagExchange == "bitstamp" {
    client = bitstamp.NewClient(flagClientId, flagApiKey, flagApiSecret)
  } else if flagExchange == "bitfinex" {
    client = bitfinex.NewClient(flagApiKey, flagApiSecret)
  } else if flagExchange == "kraken" {
    client = kraken.NewClient(flagApiKey, flagApiSecret)
  } else {
    fmt.Printf("Unknown exchange: %v\n", flagExchange)
    os.Exit(1)
  }
  client.SetDryRun(flagTest)

  openOrders, err := client.OpenOrders()
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

  A, err := client.Balance(bitcoin.FIAT)
  if err != nil {
    fmt.Printf("Error balance: %v\n", err)
    return
  }
  b, err := client.Balance(bitcoin.BTC)
  if err != nil {
    fmt.Printf("Error balance: %v\n", err)
    return
  }
  A += flagOffsetUsd
  b += flagOffsetBtc
  R := flagBtcRatio / (1 - flagBtcRatio)
  F, err := client.Fee()
  if err != nil {
    fmt.Printf("Error fee: %v\n", err)
    return
  }
  if flagSpread < 200*F {
    fmt.Printf(
      "spread (%.2f%%) must be at least twice the fee (%.2f%%) "+
        "not to make a loss.\n", flagSpread, 100*F)
    return
  }
  s := 1 + (flagSpread / 100)

  previousRate := R * A / b

  fmt.Printf("Rate = %.2f\n", previousRate)
  fmt.Printf("Fiat = %v\n", A)
  fmt.Printf("BTC = %v\n", b)
  fmt.Printf("Fee = %v\n", F)

  placeBuyOrders(A, b, R, F, s, orderMap)
  placeSellOrders(A, b, R, F, s, orderMap)

  for _, order := range orderMap {
    err := order.Execute(client)
    if err != nil {
      fmt.Printf("Error executing order: %v: %v\n", order, err)
    }
  }
}
