package main

import (
  "fmt"
  "github.com/Jeffail/gabs"
)

func main() {
  find_kids()
}

// Bad method name
func find_kids() {
  jsonParsed, err := gabs.ParseJSON([]byte(`{"object":{ "first": 1, "second": 2, "third": 3 }}`))

  if err != nil {
    panic(err)
  }

  // S is shorthand for Search
  children, _ := jsonParsed.S("object").ChildrenMap()
  for key, child := range children {
  	fmt.Printf("key: %v, value: %v\n", key, child.Data().(float64))
  }
}
