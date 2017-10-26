package main

import (
  "fmt"
  "log"
  "strings"
  "github.com/buger/jsonparser"
)

var (
  edwin = []byte(`{"person":{"name":"edwin","traits":{"relative":{"brother":[{"uno":{"pet":[{"cat":"cat-do","kittens":[{"male":"one"},{"female":"two"}]},{"dog":"dog-do"},{"rat":"rat-do"}]}}]}}}}`)
  path = "person.traits.relative.brother.[0].uno.pet.[0].kittens.[0].male"
)

func main() {

  val, ty, _, err := jsonparser.Get(edwin, fix_path(path)...)

  check_err(err)

  fmt.Printf("%s\n", val)
  fmt.Println(ty)

}

func fix_path(path string) []string {
  return strings.Split(path, ".")
}

// Break if there is an error passed in
func check_err(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
