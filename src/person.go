package main

import (
  "fmt"
  "encoding/json"
)

// Build basic person object json object
// Give the capability to custimize the object

func main() {
  template()
}

// Use this as a template

func template() {

  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
  var dat map[string]interface{}

  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  fmt.Println(dat)

  num := dat["num"].(float64)
  fmt.Println(num)

  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  fmt.Println(str1)

}
