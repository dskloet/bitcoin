package bitstamp

import (
  "fmt"
  "strconv"
)

type resultMap map[string]interface{}

func (r resultMap) getFloat(name string) float64 {
  value := r[name]
  switch value := value.(type) {
  default:
    return 0
  case float64:
    return value
  case string:
    result, err := strconv.ParseFloat(value, 64)
    if err != nil {
      fmt.Printf("Error converting: %v\n", err)
    }
    return result
  }
}

func (r resultMap) getInt(name string) int64 {
  value := r[name]
  switch value := value.(type) {
  default:
    return 0
  case int64:
    return value
  case float64:
    return int64(value)
  case string:
    result, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
      fmt.Printf("Error converting: %v\n", err)
    }
    return result
  }
}
