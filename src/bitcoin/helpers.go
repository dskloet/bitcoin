package bitcoin

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "io"
)

func JsonParse(reader io.ReadCloser, result interface{}) (err error) {
  defer reader.Close()
  buf := bytes.NewBuffer(nil)
  _, err = io.Copy(buf, reader)
  if err != nil {
    return
  }
  err = json.Unmarshal(buf.Bytes(), result)
  if err != nil {
    err = errors.New(fmt.Sprintf("Couldn't parse json: %v: %v", err, buf))
  }
  return
}
