package main

import (
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
)

func main() {
  values:= url.Values {
    "query": {"hello world"},
  }
  resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  log.Println(string(body))
}
