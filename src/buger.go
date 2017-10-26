package main

import (
  "fmt"
  "strings"
  "github.com/buger/jsonparser"
)

var (
  polly = []byte(`{"person":{"name":"polly","traits":{"relative":{"brother":[{"uno":{"pet":[{"cat":"cat-do","kittens":[{"male":"one"},{"female":"two"}]},"dog","rat"]}}]}}}}`)
  edwin = []byte(`{"person":{"name":"edwin","traits":{"relative":{"brother":[{"uno":{"pet":[{"cat":"cat-do","kittens":[{"male":"one"},{"female":"two"}]},{"dog":"dog-do"},{"rat":"rat-do"}]}}]}}}}`)
  path = "person.traits.relative.brother.[0].uno.pet"
)

func main() {

  // create_this_person("joe")
  // return

  arr := [2]string{"hello", "world"}

  fmt.Println(arr)

  val, ty, _, err := jsonparser.Get(polly, fix_path(path)...)
  check_err(err)
  fmt.Printf("%s\n", val)
  fmt.Println(ty)

  jsonparser.ArrayEach(val, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
    element, _, _, err := jsonparser.Get(value, "kittens")
    check_err(err)
    fmt.Printf("-> %s\n", element)
  })

  fmt.Println("done")

  // val, _, _, err := jsonparser.Get(edwin, fix_path(path)...)
  // check_err(err)
  // fmt.Printf("%s\n", val)

  // fmt.Println("================")
  // fmt.Printf("%s\n", edwin)
  // fmt.Println("================")
  //
  // new_val := []byte("\"dude\"")
  //
  // val, err = jsonparser.Set(edwin, new_val, fix_path(path)...)
  // check_err(err)
  // fmt.Printf("%s\n", val)

}

func create_this_person(name string) {
  raw := fmt.Sprintf(`{"person":{"name":"%s"}}`, name)
  this_person := []byte(raw)
  fmt.Printf("%s\n", this_person)
}

func fix_path(path string) []string {
  return strings.Split(path, ".")
}

// Break if there is an error passed in
// func check_err(err error) {
//   if err != nil {
//     log.Fatal(err)
//   }
// }
